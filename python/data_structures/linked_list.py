
class LinkedNode:
    data = None
    next = None

    def __init__(self, data):
        self.data = data

    def __repr__(self):
        return str(self.data)


class LinkedList:
    head = None
    tail = None

    def __init__(self, lst=None):
        lst = lst or []

        for value in lst:
            self.append(value)

    def append(self, value):
        if self.head is None:
            self.head = LinkedNode(value)
            self.tail = self.head
        else:
            self.tail.next = LinkedNode(value)
            self.tail = self.tail.next

    def search(self, value):
        node = self.head

        while node.next is not None:
            if node.data == value:
                return node

            node = node.next

    def insert_after(self, node, value):
        n = self.head

        while n.next is not None:
            if n == node:
                new_node = LinkedNode(value)
                new_node.next = n.next
                n.next = new_node
                return

            n = n.next

    def insert(self, pos, value):
        node = self.head
        index = 0

        while node:
            if index == pos - 1:
                new_node = LinkedNode(value)
                new_node.next = node.next
                node.next = new_node
                return True

            node = node.next
            index += 1
        return False

    def index(self, value):
        node = self.head
        index = 0

        while node:
            if node.data == value:
                return index

            node = node.next
            index += 1
        return None

    def __bool__(self):
        return self.head is None

    def __iter__(self):
        node = self.head

        if node is None:
            return

        while node.next is not None:
            yield node.data

            node = node.next

        yield node.data

    def __len__(self):
        counter = 0
        for _ in self:
            counter += 1

        return counter

    def __repr__(self):
        return str([x for x in self])
