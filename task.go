package tasks 

import(
	"context"
    "encoding/json"
    "fmt"
    "log"
    "time"
    "github.com/hibiken/asynq"
)

const(
	TypeEmailDelivery = "email:example"
	TypeImageresize  = "image:resize"
)

type EmailDeliveryPayload struct{
	userId int 
	templateID string
}

type ImageResizePayload struct{
	imageUrl string
}

func NewEmailDeliveryTask(userId int, templateID string) (*asynq.Task error){
    payload,err := json.Marsahl(EmailDeliveryPayload{userId:userId, templateID:templateID})
	if err != nil{
		return nil , err
	}
	return asynq.NewTask(TypeEmailDelivery, payload), nil
}

func NewImageResizeTask(src string)(*asynq.Task, error){
	payload,err := json.Marshal(ImageResizePayload({imageUrl:src}))
	if err != nil{
		return nil , err
	}
	return asynq.NewTask(TypeImageresize, payload, asynq.MaxRetry(5), asynq.Timeout(20 * time.Minute)),nil
}

func HandleEmailDeliveryTask(ctx,context.Context, t.*asynq.Task) error {
	var p EmailDeliveryPayload 
	if err := json.Unmasrhal(t.Payload(), &p);err != nil {
        return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
    }
	log.Printf("Sending Email to User: user_id=%d, template_id=%s", p.UserID, p.TemplateID)
    // Email delivery code ...
    return nil
}
type ImageProcessor struct {
    // ... fields for struct
}

func (processor *ImageProcessor) ProcessTask(ctx context.Context, t.*asyqn.Task) error{
	var p ImageResizePayload 
	if err := json.Unmarshal(t.Payload(), &p);err != nil {
        return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
    }
    log.Printf("Resizing image: src=%s", p.SourceURL)
    // Image resizing code ...
    return nil
}

func NewImageProcessor() *ImageProcessor {
	return &ImageProcessor{}
}