from data_structures.node import Node


class Queue:
    head = None
    tail = None

    def __init__(self, lst=None):
        lst = lst or []

        for value in lst:
            self.push(value)

    def push(self, value):
        new_node = Node(value)

        if self.head is None:
            self.head = new_node
            self.tail = self.head
        else:
            self.tail.next = new_node
            self.tail = new_node

    def pop(self):
        if self.head:
            data = self.head.data
            self.head = self.head.next
            return data

    def peek(self):
        return self.head.data
