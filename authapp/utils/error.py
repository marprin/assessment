class DuplicateValueError(Exception):
    def __init__(self, message: str = "Duplicate value") -> None:
        self.message = message
        super().__init__(self.message)

    def __str__(self) -> str:
        return self.message
