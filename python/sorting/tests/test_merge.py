import random

from unittest import TestCase

from sorting.merge import sort


class TestMergeSort(TestCase):

    def test_sort(self):
        self.assertEqual(sort([6, 8, 45, 1, 84, 149, 9, 17]), [1, 6, 8, 9, 17, 45, 84, 149])

        l = []
        self.assertEqual(sort(l), sorted(l))

        l = [random.randint(0, 100) for _ in range(10)]
        self.assertEqual(sort(l), sorted(l))

        l = [random.randint(0, 100) for _ in range(25)]
        self.assertEqual(sort(l), sorted(l))

        l = [random.randint(5, 250) for _ in range(150)]
        self.assertEqual(sort(l), sorted(l))

        l = [random.randint(0, 500) for _ in range(1500)]
        self.assertEqual(sort(l), sorted(l))
