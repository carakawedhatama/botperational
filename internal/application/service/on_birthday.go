package service

import (
	"context"
	"math/rand"
	"time"

	"botperational/config"
	"botperational/internal/domain/on_birthday"
	"botperational/internal/pkg/discord"
)

type OnBirthday interface {
	ProcessOnBirthdayData(ctx context.Context) error
}

type OnBirthdayService struct {
	OnBirthdayRepo on_birthday.Repository `inject:"onBirthdayRepository"`
	Cfg            *config.Config         `inject:"config"`
}

func (s *OnBirthdayService) ProcessOnBirthdayData(ctx context.Context) error {
	emp, err := s.getOnBirthdayData(ctx)
	if err != nil {
		return err
	}

	if len(emp) > 0 {

		contents := s.prepareContentData(emp)

		if len(contents) > 0 {
			for _, content := range contents {
				err = discord.SentViaWebhook(on_birthday.DOC_TYPE, s.Cfg.Discord.WebhookUrl.OnBirthday, content)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (s *OnBirthdayService) getOnBirthdayData(ctx context.Context) ([]*on_birthday.OnBirthday, error) {
	data, err := s.OnBirthdayRepo.GetOnBirthdayEmployee(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *OnBirthdayService) prepareContentData(data []*on_birthday.OnBirthday) []*discord.Discord {
	min := 0
	max := len(on_birthday.MessagesIndex)
	max2 := len(on_birthday.FooterIndex)

	contents := make([]*discord.Discord, 0)

	for _, x := range data {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		randomIndex := r.Intn(max-1) + min

		contentMsg := on_birthday.BirthdayContent(randomIndex)
		description := on_birthday.BirthdayDescription(x.EmpName, x.DeptName, x.PosName, x.Gender, x.Age)
		img := x.BirthdayPosterUrl

		randomIndex2 := r.Intn(max2-1) + min
		wishes := on_birthday.BirthdayWishes(randomIndex2)

		footer := discord.DiscordFooters{
			Text: wishes,
		}

		image := discord.DiscordImages{
			Url: img,
		}

		embed := discord.DiscordEmbeds{
			Description: description,
			Image:       image,
			Footer:      footer,
		}
		embeds := make([]discord.DiscordEmbeds, 0)

		embeds = append(embeds, embed)

		content := &discord.Discord{
			AvatarUrl: s.Cfg.Discord.AvatarUrl.BirthdayAvatar,
			Content:   contentMsg,
			Embeds:    embeds,
		}

		contents = append(contents, content)

		duration := time.Duration(2) * time.Millisecond
		time.Sleep(duration)
	}

	return contents
}
