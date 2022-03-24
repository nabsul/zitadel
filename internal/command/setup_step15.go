package command

import (
	"context"

	"github.com/caos/logging"
	"github.com/caos/zitadel/internal/domain"
	"github.com/caos/zitadel/internal/eventstore"
)

type Step15 struct {
	DefaultMailTemplate domain.MailTemplate
}

func (s *Step15) Step() domain.Step {
	return domain.Step15
}

func (s *Step15) execute(ctx context.Context, commandSide *Commands) error {
	return commandSide.SetupStep15(ctx, s)
}

func (c *Commands) SetupStep15(ctx context.Context, step *Step15) error {
	fn := func(iam *InstanceWriteModel) ([]eventstore.Command, error) {
		_, mailTemplateEvent, err := c.changeDefaultMailTemplate(ctx, &step.DefaultMailTemplate)
		if err != nil {
			return nil, err
		}
		logging.Log("SETUP-2nfsd").Info("default mail template/text set up")
		return []eventstore.Command{mailTemplateEvent}, nil
	}
	return c.setup(ctx, step, fn)
}
