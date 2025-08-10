package order

import (
	"context"
	"log"
	"time"

	"github.com/PrajwalG12121998/E-commerce-microservice-application-using-Golang/order/pb"
	"google.golang.org/grpc"
)

// Client represents a client for the order service
type Client struct {
	conn    *grpc.ClientConn
	service pb.OrderServiceClient
}

// NewClient creates a new order service client
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := pb.NewOrderServiceClient(conn)
	return &Client{
		conn:    conn,
		service: client,
	}, nil
}

// Close closes the client connection
func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) PostOrder(ctx context.Context, accountID string, products []OrderedProduct) (*Order, error) {
	protoProducts := []*pb.PostOrderRequest_OrderProduct{}
	for _, p := range products {
		protoProducts = append(protoProducts, &pb.PostOrderRequest_OrderProduct{
			ProductId: p.ID,
			Quantity:  p.Quantity,
		})
	}
	r, err := c.service.PostOrder(
		ctx,
		&pb.PostOrderRequest{
			AccountId: accountID,
			Products:  protoProducts,
		},
	)
	if err != nil {
		return nil, err
	}
	newOrder := r.Order
	newOrderCreatedAt := time.Time{}
	newOrderCreatedAt.UnmarshalBinary(newOrder.CreatedAt)
	return &Order{
		ID:         newOrder.Id,
		CreatedAt:  newOrderCreatedAt,
		TotalPrice: newOrder.TotalPrice,
		AccountID:  newOrder.AccountId,
		Products:   products,
	}, nil
}

func (c *Client) GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error) {
	log.Println("Inside client.go in Order microservice")

	r, err := c.service.GetOrdersForAccount(ctx, &pb.GetOrdersForAccountRequest{
		AccountId: accountID,
	})
	log.Println("printing the response from the server")
	log.Println(r)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	orders := []Order{}
	for _, orderProto := range r.Orders {
		newOrder := Order{
			ID:         orderProto.Id,
			TotalPrice: orderProto.TotalPrice,
			AccountID:  orderProto.AccountId,
		}
		newOrder.CreatedAt = time.Time{}
		newOrder.CreatedAt.UnmarshalBinary(orderProto.CreatedAt)
		products := []OrderedProduct{}
		for _, p := range orderProto.Products {
			log.Println(newOrder.Products)
			products = append(products, OrderedProduct{
				ID:          p.Id,
				Quantity:    p.Quantity,
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
			})
		}
		newOrder.Products = products
		orders = append(orders, newOrder)
	}
	log.Println("Printing logs in client.go of orders")
	log.Println(orders)
	return orders, nil
}
