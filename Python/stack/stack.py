class Node:
    def __init__(self, value):
        self.value = value
        self.next = None

class Stack:
    def __init__(self):
        self.top = None
        self.height = 0

    # push a value to the top of the stack
    def push(self, value):
        new_node = Node(value)
        new_node.next = self.top
        self.top = new_node
        self.height += 1
    
    # pop the top value from the stack
    def pop(self):
        if self.height == 0:
            return None
        value = self.top.value
        self.top = self.top.next
        self.height -= 1 
        return value
    
    # peek the top value from the stack
    def peek(self):
        if self.height ==0:
            return None
        return self.top.value

    # print the stack
    def print_stack(self):
        temp = self.top
        print("Stack:")
        while temp is not None:
            print(temp.value)
            temp = temp.next
        print("End of Stack")