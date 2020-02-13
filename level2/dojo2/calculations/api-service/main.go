package main
import (
	pb "calculations/pb/fib/v1"
	piService "calculations/pb/pi/v1"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)
func main() {

	type FibJSON struct {
		K int64 `json:"k"`
	}
	/*---FI Service---*/
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	fibClient := pb.NewFibServiceClient(conn)
	/*---PI Service---*/
	conn_pi, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	piClient := piService.NewPiServiceClient(conn_pi)

	routes := mux.NewRouter()
	routes.HandleFunc("/", indexHandler).Methods("GET")
	routes.HandleFunc("/fibonacci", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var fibJSON FibJSON
		err := json.NewDecoder(r.Body).Decode(&fibJSON)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
		defer cancel()
		req := &pb.FibRequest{FibNum: uint64(fibJSON.K)}
		if resp, err := fibClient.Compute(ctx, req); err == nil {
			fmt.Println("Fibonacci number is", resp.Result)
			json.NewEncoder(w).Encode(map[string]uint64{"fibonacci_number": resp.Result})
		} else {
			msg := fmt.Sprintf("Internal server error: %s", err.Error())
			fmt.Println("Internal server error:", err.Error())
			json.NewEncoder(w).Encode(msg)
		}
	}).Methods("POST")
	routes.HandleFunc("/pi", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
		defer cancel()
		req := &piService.Empty{}
		if resp, err := piClient.Compute(ctx, req); err == nil {
			fmt.Println("Pi number is", resp.Result)
			json.NewEncoder(w).Encode(map[string]float32{"pi": resp.Result})
		} else {
			msg := fmt.Sprintf("Internal server error: %s", err.Error())
			fmt.Println("Internal server error:", err.Error())
			json.NewEncoder(w).Encode(msg)
		}
	}).Methods("GET")
	fmt.Println("Application is running on :8080")
	http.ListenAndServe(":8080", routes)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UFT-8")
	json.NewEncoder(w).Encode("Server is running")
}