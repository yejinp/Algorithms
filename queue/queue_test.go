package queue

import "testing"

func TestStackEmpty(t *testing.T) {
	err, q := QueueInit(1)
	if err == nil {
		t.Errorf("s.IsEmpty() got unnil, want nil")
	}

	err, q = QueueInit(4)
	if err != nil {
		t.Errorf("s.IsEmpty() got nil, want unnil")
	}

	if !q.IsEmpty() {
		t.Errorf("s.IsEmpty() got %v, want: %v", false, true)
	}
}

func TestStackEnqueue(t *testing.T) {
	err, q := QueueInit(4)
	if err != nil {
		t.Errorf("s.IsEmpty() got nil, want unnil")
	}

	if err = q.Enqueue(1); err != nil {
		t.Errorf("s.Enqueue() got %s, want nil", err)
	}

	if err, _ = q.Dequeue(); err != nil {
		t.Errorf("s.Dequeue() got %v, want unnil", err)
	}

}

func TestStackDequeue(t *testing.T) {
	err, q := QueueInit(4)
	if err != nil {
		t.Errorf("s.QueueInit() got nil, want unnil")
	}

	if err, _ = q.Dequeue(); err == nil {
		t.Errorf("s.Dequeue() got nil, want unnil")
	}

	if err = q.Enqueue(1); err != nil {
		t.Errorf("s.Enqueue() got %s, want nil", err)
	}

	if err = q.Enqueue(2); err != nil {
		t.Errorf("s.Enqueue() got %s, want nil", err)
	}

	if err = q.Enqueue(3); err != nil {
		t.Errorf("s.Enqueue() got %s, want nil", err)
	}

	if err = q.Enqueue(4); err == nil {
		t.Errorf("s.Enqueue() got %s, want unnil", err)
	}

}
