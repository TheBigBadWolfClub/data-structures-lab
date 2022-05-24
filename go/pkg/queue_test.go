package pkg

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	t.Parallel()

	if got := NewQueue[int](); !reflect.DeepEqual(got, &queue[int]{}) {
		t.Errorf("NewQueue() fail to create int Queue")
	}

	if got := NewQueue[string](); !reflect.DeepEqual(got, &queue[string]{}) {
		t.Errorf("NewQueue() fail to create string Queue")
	}

}

func Test_queue_Enqueue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		elem    int
		initial Queue[int]
		want    Queue[int]
	}{
		{
			name:    "add to empty queue",
			elem:    0,
			initial: &queue[int]{},
			want:    &queue[int]{mem: []int{0}},
		}, {
			name:    "add to queue with one elem",
			elem:    1,
			initial: &queue[int]{mem: []int{0}},
			want:    &queue[int]{mem: []int{0, 1}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.initial.Enqueue(tt.elem)

			if diff := cmp.Diff(tt.initial, tt.want, cmp.AllowUnexported(queue[int]{})); diff != "" {
				t.Errorf("fail to Enqueue diff=%s", diff)
			}
		})
	}
}

func Test_queue_Dequeue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		initial  Queue[int]
		wantOK   bool
		wantElem int
		want     Queue[int]
	}{
		{
			name:     "remove from empty queue",
			initial:  &queue[int]{},
			wantOK:   false,
			wantElem: 0,
			want:     &queue[int]{},
		}, {
			name:     "remove from queue with one elem",
			initial:  &queue[int]{mem: []int{0}},
			wantOK:   true,
			wantElem: 0,
			want:     &queue[int]{mem: []int{}},
		}, {
			name:     "remove from queue with multiple elem",
			initial:  &queue[int]{mem: []int{0, 1, 2}},
			wantOK:   true,
			wantElem: 0,
			want:     &queue[int]{mem: []int{1, 2}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			elem, ok := tt.initial.Dequeue()
			if elem != tt.wantElem || ok != tt.wantOK {
				t.Fatalf("fail to Dequeue wantElem=%v,%v, got=%v,%v", tt.wantElem, tt.wantOK, elem, ok)
			}

			if diff := cmp.Diff(tt.initial, tt.want, cmp.AllowUnexported(queue[int]{})); diff != "" {
				t.Fatalf("fail to Dequeue diff=%s", diff)
			}
		})
	}
}

func Test_queue_Peek(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		initial  Queue[int]
		wantOK   bool
		wantElem int
		want     Queue[int]
	}{
		{
			name:     "remove from empty queue",
			initial:  &queue[int]{},
			wantOK:   false,
			wantElem: 0,
			want:     &queue[int]{},
		}, {
			name:     "remove from queue with one elem",
			initial:  &queue[int]{mem: []int{0}},
			wantOK:   true,
			wantElem: 0,
			want:     &queue[int]{mem: []int{0}},
		}, {
			name:     "remove from queue with multiple elem",
			initial:  &queue[int]{mem: []int{0, 1, 2}},
			wantOK:   true,
			wantElem: 0,
			want:     &queue[int]{mem: []int{0, 1, 2}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			elem, ok := tt.initial.Peek()
			if elem != tt.wantElem || ok != tt.wantOK {
				t.Fatalf("fail to Peek wantElem=%v,%v, got=%v,%v", tt.wantElem, tt.wantOK, elem, ok)
			}

			if diff := cmp.Diff(tt.initial, tt.want, cmp.AllowUnexported(queue[int]{})); diff != "" {
				t.Fatalf("fail to Peek diff=%s", diff)
			}
		})
	}
}

func Test_queue_IsFull(t *testing.T) {
	t.Parallel()
	t.SkipNow()
}

func Test_queue_IsEmpty(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		initial Queue[int]
		want    bool
	}{
		{
			name:    "is empty",
			initial: &queue[int]{},
			want:    true,
		}, {
			name:    "not empty",
			initial: &queue[int]{mem: []int{0}},
			want:    false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.initial.IsEmpty()
			if tt.want != got {
				t.Errorf("fail to IsEmpty want=%v, got=%v", tt.want, got)
			}
		})
	}
}
