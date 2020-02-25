package queue
import "testing"

func TestEnqueue(t *testing.T){
	q := NewQueue(3)
	if err := q.Enqueue("a");err!= nil{
		t.Fatal("Error in enqueue: ",err)
	}
	q.Enqueue("b")
	if q.PeekFront() != "a"{
		t.Errorf("Enqueue not working properly: PeekFront: %s",q.PeekFront())
	}
	q.Enqueue("c")
	if err := q.Enqueue("d");err != ErrQueueFull{
		t.Errorf("Queue is full but error not popped up: %s",err.Error())
	}
}

func TestDequeue(t *testing.T){
	q := NewQueue(2)
	q.Enqueue("a")
	q.Enqueue("b")
	if err := q.Dequeue();err != nil{
		t.Errorf("Error in dequeue: %s",err)
	}
	if q.PeekFront() != q.PeekRear(){
		t.Errorf("Error in peeking after dequeue: %s",q.PeekFront())
	}
}