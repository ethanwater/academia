#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>

typedef struct Node {
    int value;
    struct Node *next;
} Node;

typedef struct Queue {
    Node *head;
    Node *tail;
    pthread_mutex_t mux; 
} Queue;

Node* create_node(int value) {
    Node *node = (Node*)malloc(sizeof(Node));
    if (node != NULL) {
        node->value = value;
        node->next = NULL;
    }
    return node;
}

void init(Queue *queue) {
    queue->head = NULL;
    queue->tail = NULL;
    pthread_mutex_init(&queue->mux, NULL); 
}

void enqueue(Queue *queue, int value) {
    Node *new_node = create_node(value);
    if (new_node != NULL) {
        pthread_mutex_lock(&queue->mux); 
        if (queue->head == NULL) {
            queue->head = new_node;
        }
        if (queue->tail != NULL) {
            queue->tail->next = new_node;
        }
        queue->tail = new_node;
        pthread_mutex_unlock(&queue->mux); 
    }
}

int dequeue(Queue *queue) {
    int value = -1;

    pthread_mutex_lock(&queue->mux); 
    if (queue->head != NULL) {
        Node *temp = queue->head;
        value = temp->value;
        queue->head = temp->next;
        free(temp);
    }
    pthread_mutex_unlock(&queue->mux); 
    return value;
}

void test_enqueue(Queue *queue) {
    printf("Head value: %d\n", queue->head ? queue->head->value : 0);
    enqueue(queue, 1);
    printf("Head value after enqueue(1): %d\n", queue->head->value);
    printf("Tail value: %d\n", queue->tail->value);
    enqueue(queue, 2);
    printf("Head value after enqueue(2): %d\n", queue->head->value);
    printf("Tail value: %d\n", queue->tail->value);
    enqueue(queue, 3);
    printf("Head value after enqueue(3): %d\n", queue->head->value);
    printf("Tail value: %d\n", queue->tail->value);
}

void* threader(void *arg) {
    Queue *queue = (Queue*)arg;
    test_enqueue(queue);
    pthread_exit(NULL);
}

int main() {
    Queue queue;
    init(&queue);

    pthread_t thread;
    pthread_create(&thread, NULL, threader, &queue);

    test_enqueue(&queue);
    pthread_join(thread, NULL);
    pthread_mutex_destroy(&queue.mux); 
    return 0;
}
