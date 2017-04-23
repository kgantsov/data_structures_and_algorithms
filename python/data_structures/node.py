
class Node:
    data = None
    next = None

    def __init__(self, data):
        self.data = data

    def __repr__(self):
        return str(self.data)
