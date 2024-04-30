package service

import (
	"context"
	"fmt"
	"strings"

	"botperational/config"
	"botperational/internal/domain/on_leave"
	"botperational/internal/pkg/discord"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type OnLeave interface {
	ProcessOnLeaveData(ctx context.Context) error
}

type OnLeaveService struct {
	OnLeaveRepo on_leave.Repository `inject:"onLeaveRepository"`
	Cfg         *config.Config      `inject:"config"`
}

func (s *OnLeaveService) getOnLeaveData(ctx context.Context) ([]*on_leave.OnLeave, error) {
	data, err := s.OnLeaveRepo.GetOnLeaveEmployee(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *OnLeaveService) prepareContentData(data []*on_leave.OnLeave) (*discord.Discord, error) {
	var (
		footerText  = on_leave.FOOTER_TEXT
		contentMsg  = on_leave.ContentMsg()
		description = ""
	)

	for _, x := range data {
		for _, c := range []cases.Caser{
			cases.Title(language.Indonesian),
		} {
			description += fmt.Sprintf("**%s** (%s) \n", c.String(strings.ToLower(x.EmpName)), x.DeptName)
		}
	}

	footers := discord.DiscordFooters{
		Text: footerText,
	}

	embed := discord.DiscordEmbeds{
		Description: description,
		Footer:      footers,
	}

	embeds := make([]discord.DiscordEmbeds, 0)
	embeds = append(embeds, embed)

	discord := &discord.Discord{
		AvatarUrl: s.Cfg.Discord.AvatarUrl.LeaveAvatar,
		Content:   contentMsg,
		Embeds:    embeds,
	}

	return discord, nil
}

func (s *OnLeaveService) ProcessOnLeaveData(ctx context.Context) error {
	emp, err := s.getOnLeaveData(ctx)
	if err != nil {
		return err
	}

	if len(emp) > 0 {

		content, err := s.prepareContentData(emp)
		if err != nil {
			return err
		}

		err = discord.SentViaWebhook(on_leave.DOC_TYPE, s.Cfg.Discord.WebhookUrl.OnLeave, content)
		if err != nil {
			return err
		}
	} else {
		content := &discord.Discord{
			AvatarUrl: s.Cfg.Discord.AvatarUrl.WorkAvatar,
			Content:   on_leave.NoLeaveData(),
		}

		err = discord.SentViaWebhook(on_leave.DOC_TYPE, s.Cfg.Discord.WebhookUrl.OnLeave, content)
		if err != nil {
			return err
		}
	}

	return nil
}
