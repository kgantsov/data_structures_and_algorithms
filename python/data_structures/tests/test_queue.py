from unittest import TestCase

from data_structures.queue import Queue


class TestQueue(TestCase):

    def test_init(self):
        queue = Queue([6, 8, 45, 1, 84, 149, 9, 17])

        self.assertEqual(queue.pop(), 6)
        self.assertEqual(queue.pop(), 8)
        self.assertEqual(queue.pop(), 45)
        self.assertEqual(queue.pop(), 1)
        self.assertEqual(queue.pop(), 84)
        self.assertEqual(queue.pop(), 149)
        self.assertEqual(queue.pop(), 9)
        self.assertEqual(queue.pop(), 17)

    def test_push(self):
        queue = Queue()

        queue.push(1)
        self.assertEqual(queue.peek(), 1)

        queue.push(50)
        self.assertEqual(queue.peek(), 1)

        queue.push(30)
        self.assertEqual(queue.peek(), 1)

        queue.push(70)
        self.assertEqual(queue.peek(), 1)

    def test_pop(self):
        queue = Queue()
        queue.push(1)
        queue.push(50)
        queue.push(30)
        queue.push(70)

        self.assertEqual(queue.pop(), 1)
        self.assertEqual(queue.pop(), 50)

        queue.push(75)
        self.assertEqual(queue.peek(), 30)

        self.assertEqual(queue.pop(), 30)
        self.assertEqual(queue.pop(), 70)

        queue.push(876)
        self.assertEqual(queue.peek(), 75)

        self.assertEqual(queue.pop(), 75)
        self.assertEqual(queue.pop(), 876)
