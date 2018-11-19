package stub

import (
	"context"

	"github.com/philbrookes/finalizer-operator/pkg/apis/philbrookes/finalizer"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	if event.Deleted {
		return nil
	}
	switch o := event.Object.(type) {
	case *finalizer.Item:
		o, err := handleItem(o)
		if err != nil {
			return err
		}
		return sdk.Update(o)

	}
	return nil
}

func handleItem(i *finalizer.Item) (*finalizer.Item, error) {
	if i.DeletionTimestamp != nil {
		logrus.Infof("removing finalizer...")
		return removeFinalizer(i)
	}
	logrus.Infof("adding finalizer...")
	return addFinalizer(i)
}

func removeFinalizer(cr *finalizer.Item) (*finalizer.Item, error) {
	for i, v := range cr.Finalizers {
		if v == "phil.brookes/finalizer" {
			cr.Finalizers = append(cr.Finalizers[:i], cr.Finalizers[i+1:]...)
		}
	}
	return cr, nil
}

func addFinalizer(cr *finalizer.Item) (*finalizer.Item, error) {
	for _, v := range cr.Finalizers {
		logrus.Infof("finalizer found: %v", v)
		//already exists, return unmodified
		if v == "phil.brookes/finalizer" {
			return cr, nil
		}
	}
	cr.Finalizers = append(cr.Finalizers, "phil.brookes/finalizer")
	return cr, nil
}
