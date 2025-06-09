class Node:
    def __init__(self, value):
        self.value = value
        self.next = None

    def __str__(self):
        return str(f"Node: {self.value}")


class LinkedList:
    def __init__(self, values=None):
        self.head = None
        self.tail = None
        self.length = 0
        if values is not None:
            for value in values:
                self.append(value)

    # print the linked list
    def __str__(self):
        values = []
        current_node = self.head
        while current_node is not None:
            values.append(current_node.value)
            current_node = current_node.next
        return f"LinkedList: {values}"

    # append a value to the end of the linked list
    def append(self, value):
        new_node = Node(value)
        if self.head is None:
            self.head = new_node
            self.tail = new_node
        else:
            self.tail.next = new_node
            self.tail = new_node
        self.length += 1

    # prepend a value to the beginning of the linked list
    def prepend(self, value):
        new_node = Node(value)
        if self.head is None:
            self.head = new_node
            self.tail = new_node
        else:
            new_node.next = self.head
            self.head = new_node
        self.length += 1

    # pop the last value of the linked list
    def pop(self):
        if self.length == 0:
            return None
        temp = self.head
        if self.length == 1:
            value = self.head.value
            self.head = None
            self.tail = None
        else:
            temp = self.head
            while temp.next is not self.tail:
                temp = temp.next
            value = self.tail.value
            self.tail = temp
            self.tail.next = None
        self.length -= 1
        return value

    # pop the first value of the linked list
    def pop_first(self):
        if self.length == 0:
            return None
        if self.length == 1:
            value = self.head.value
            self.head = None
            self.tail = None
        else:
            value = self.head.value
            self.head = self.head.next
        self.length -= 1
        return value

    # get the node at the given index
    def get(self, index):
        if index < 0 or index >= self.length:
            return None
        temp = self.head
        for _ in range(index):
            temp = temp.next
        return temp

    # set the value of the node at the given index
    def set(self, index, value):
        temp = self.get(index)
        if temp is not None:
            temp.value = value
            return True
        return False

    # insert a value at the given index
    def insert(self, index, value):
        if index < 0 or index > self.length:
            return False
        if index == 0:
            return self.prepend(value)
        if index == self.length:
            return self.append(value)
        temp = self.get(index - 1)
        if temp is not None:
            new_node = Node(value)
            new_node.next = temp.next
            temp.next = new_node
            self.length += 1
            return True
        return False

    # remove the node at the given index
    def remove(self, index):
        if index < 0 or index >= self.length:
            return None
        if index == 0:
            return self.pop_first()
        if index == self.length - 1:
            return self.pop()
        temp = self.get(index - 1)
        if temp is not None:
            temp.next = temp.next.next
            self.length -= 1
            return True
        return False

    # reverse the linked list
    def reverse(self):
        if self.length == 0 or self.length == 1:
            return
        current = self.head
        prev = None
        while current is not None:
            next = current.next
            current.next = prev
            prev = current
            current = next
        self.head, self.tail = self.tail, self.head