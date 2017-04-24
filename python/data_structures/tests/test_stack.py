from unittest import TestCase

from data_structures.stack import Stack


class TestStack(TestCase):

    def test_init(self):
        stack = Stack([6, 8, 45, 1, 84, 149, 9, 17])

        self.assertEqual(stack.pop(), 17)
        self.assertEqual(stack.pop(), 9)
        self.assertEqual(stack.pop(), 149)
        self.assertEqual(stack.pop(), 84)
        self.assertEqual(stack.pop(), 1)
        self.assertEqual(stack.pop(), 45)
        self.assertEqual(stack.pop(), 8)
        self.assertEqual(stack.pop(), 6)

    def test_push(self):
        stack = Stack()

        stack.push(1)
        self.assertEqual(stack.peek(), 1)

        stack.push(50)
        self.assertEqual(stack.peek(), 50)

        stack.push(30)
        self.assertEqual(stack.peek(), 30)

        stack.push(70)
        self.assertEqual(stack.peek(), 70)

    def test_pop(self):
        stack = Stack()
        stack.push(1)
        stack.push(50)
        stack.push(30)
        stack.push(70)

        self.assertEqual(stack.pop(), 70)
        self.assertEqual(stack.pop(), 30)

        stack.push(75)
        self.assertEqual(stack.peek(), 75)

        self.assertEqual(stack.pop(), 75)
        self.assertEqual(stack.pop(), 50)

        stack.push(876)
        self.assertEqual(stack.peek(), 876)

        self.assertEqual(stack.pop(), 876)
        self.assertEqual(stack.pop(), 1)
