from data_structures.node import Node


class Stack:
    top = None

    def push(self, value):
        new_top = Node(value)

        if self.top is None:
            self.top = new_top
        else:
            new_top.next = self.top
            self.top = new_top

    def pop(self):
        if self.top:
            node = self.top
            self.top = node.next
            return node.data

    def peek(self):
        return self.top.data
