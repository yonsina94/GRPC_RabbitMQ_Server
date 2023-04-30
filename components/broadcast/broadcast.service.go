package broadcast

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/yonsina94/GRPC_RabbitMQ_Server/clients/ampq"
	"github.com/yonsina94/GRPC_RabbitMQ_Server/pkg/pb"
	"google.golang.org/protobuf/encoding/protojson"
)

type BroadcastService struct {
	cli *ampq.Client
	pb.UnimplementedBroadcasterServiceServer
}

func (svc *BroadcastService) SendMessage(payload *pb.Payload, srv pb.BroadcasterService_SendMessageServer) error {
	log.Printf("Propagating message for %s",payload.Action)

	var amount int = 10;

	log.Printf("Amount of messages to send: %d",amount)

	var wg sync.WaitGroup

	for i := 0; i < amount; i++ {
		wg.Add(1)

		go func (count int64)  {
			defer wg.Done()

			var (
				data json.RawMessage
				resp pb.SendResult
			)

			data, err := protojson.Marshal(payload.Data)
			if err != nil {
				e := err.Error()
				resp = pb.SendResult {Identifier: "", Sended: false, ErrorMessage: &e}

                log.Printf("Error marshalling data: %v",err)
            }


			err = svc.cli.Publish(payload.Action,data)
			if err != nil {
				e := err.Error()
				resp = pb.SendResult {Identifier: "", Sended: false, ErrorMessage: &e}

				log.Printf("Error publishing message: %v",err)
			}

			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v",err)
			} else {
				log.Printf("Message sended to this conected exchange. Number of exchange: %d", count)
			}
		}(int64(i))
	}

	wg.Wait()

	return nil
}