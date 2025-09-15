from collections import deque

def auto_operations(cls):
    """ Forward arithmetic operations to attribute _value """
    ops = {
        '__add__': lambda self, other: self._value + other,
        '__radd__': lambda self, other: other + self._value,
        '__sub__': lambda self, other: self._value - other,
        '__rsub__': lambda self, other: other - self._value,
        '__mul__': lambda self, other: self._value * other,
        '__rmul__': lambda self, other: other * self._value,
        '__gt__': lambda self, other: self._value > other,
        '__lt__': lambda self, other: self._value < other
  #      '__eq__': lambda self, other : self._value == other
    }
    for name, func in ops.items():
        setattr(cls, name, func)
    return cls

@auto_operations
class InputCell:
    """ Cells with settable values """
    def __init__(self, initial_value):
        self.superiors = []
        self._value = initial_value

    def set_superior(self, superior):
        self.superiors.append(superior)

    def __setattr__(self, name, value):
        if name != "value":
            object.__setattr__(self, name, value)
        else:
            self._value = value
            # BFS-like walkthrough to avoid eager recursion
            to_visit = deque(self.superiors[:])
            visited = set()
            visited.add(self)


            while len(to_visit) > 0:
                cell = to_visit.popleft()
                if cell in visited :
                    continue
                if any(i in to_visit for i in cell.inputs):
                    to_visit.append(cell) # Waiting for all its dependencies to be computed first
                else : 
                    visited.add(cell)
                    cell.recompute()
                    to_visit.extend(s for s in cell.superiors if s not in visited)

    def __getattribute__(self, name: str):
        if name == "value" : return self._value
        else : return object.__getattribute__(self, name)

class ComputeCell(InputCell):
    """ Cells computed in terms of other cells """
    def __init__(self, inputs, compute_function):
        super().__init__(compute_function(inputs))
        self.compute_f = compute_function
        self.inputs = inputs
        self.callbacks = []
        for i in inputs:
            i.set_superior(self)

    def recompute(self):
        new_value = self.compute_f(self.inputs)
        if self._value != new_value:
            self._value = new_value
            for c in self.callbacks :
                c(self._value)

    def add_callback(self, callback):
        self.callbacks.append(callback)

    def remove_callback(self, callback):
        if callback in self.callbacks:
            self.callbacks.remove(callback)