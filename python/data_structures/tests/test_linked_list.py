import random
from unittest import TestCase

from data_structures.linked_list import LinkedList


class TestLinkedList(TestCase):

    def test_init(self):
        lst = LinkedList()
        self.assertEqual(list(lst), [])

        lst = LinkedList([])
        self.assertEqual(list(lst), [])

        lst = LinkedList([2, 5,  7, 10])
        self.assertEqual(list(lst), [2, 5, 7, 10])

        expected_list = [random.randint(0, 100) for _ in range(20)]
        lst = LinkedList(expected_list)
        self.assertEqual(list(lst), expected_list)

    def test_append(self):

        lst = LinkedList([82, 57, 16, 20, 21, 84, 99, 56, 100, 46])

        lst.append(88)
        lst.append(99)
        lst.append(36)
        lst.append(45)

        self.assertEqual(list(lst), [82, 57, 16, 20, 21, 84, 99, 56, 100, 46, 88, 99, 36, 45])

    def test_length(self):
        lst = LinkedList()
        self.assertEqual(len(lst), 0)

        lst = LinkedList([random.randint(0, 100) for _ in range(20)])
        self.assertEqual(len(lst), 20)

        lst = LinkedList([random.randint(0, 100) for _ in range(48)])
        self.assertEqual(len(lst), 48)

        lst = LinkedList([random.randint(0, 100) for _ in range(7)])
        self.assertEqual(len(lst), 7)

        lst = LinkedList([random.randint(0, 100) for _ in range(33)])
        self.assertEqual(len(lst), 33)

    def test_empty(self):
        self.assertFalse(bool(LinkedList()))
        self.assertFalse(bool(LinkedList([])))
        self.assertTrue(bool(LinkedList([0])))
        self.assertTrue(bool(LinkedList([1, 4, 5])))

    def test_search(self):
        lst = LinkedList([82, 57, 16, 20, 21, 84, 99, 56, 100, 46])

        node = lst.search(99)
        self.assertEqual(node.data, 99)

        node = lst.search(57)
        self.assertEqual(node.data, 57)

        node = lst.search(777)
        self.assertIsNone(node)

    def test_insert_after(self):
        lst = LinkedList([82, 57, 16, 20, 21, 84, 99, 56, 100, 46])

        node = lst.search(99)
        lst.insert_after(node, 555)

        node = lst.search(57)
        lst.insert_after(node, 86)

        self.assertEqual(list(lst), [82, 57, 86, 16, 20, 21, 84, 99, 555, 56, 100, 46])

    def test_insert(self):
        lst = LinkedList([82, 57, 16, 20, 21, 84, 99, 56, 100, 46])

        inserted = lst.insert(7, 555)
        self.assertTrue(inserted)

        inserted = lst.insert(2, 86)
        self.assertTrue(inserted)

        inserted = lst.insert(200, 86)
        self.assertFalse(inserted)

        self.assertEqual(list(lst), [82, 57, 86, 16, 20, 21, 84, 99, 555, 56, 100, 46])

    def test_index(self):
        lst = LinkedList([82, 57, 16, 20, 21, 84, 99, 56, 100, 46])

        index = lst.index(99)
        self.assertEqual(index, 6)

        index = lst.index(57)
        self.assertEqual(index, 1)

        index = lst.index(46)
        self.assertEqual(index, 9)

        index = lst.index(777)
        self.assertIsNone(index)
