"""Exceptions classes module."""


class UnsupportedFileFormatError(Exception):
    """Unsupported file extension in --file parameter."""

    def __init__(self, *args):
        self.message: str = "Unsupported file format."
        super().__init__(*args)


class WrongMazeDimensionsError(Exception):
    """Maze dimensions in wrong format."""

    def __init__(self, *args):
        self.message: str = "Maze dimensions wrong format."
        super().__init__(*args)


class UnprovidedMazeDataError(Exception):
    """No data to build maze or get an existing."""

    def __init__(self, *args):
        self.message: str = "No maze parameters or an existing one were given."
        super().__init__(*args)

class WrongSolvingIndiciesError(Exception):
    """Wrong solving indicies provided."""
    
    def __init__(self, *args):
        self.message: str = "Solving indicies were provided in wrong format."
        super().__init__(*args)

class MazeFileNotFoundError(Exception):
    """No file could be found by a given path."""
    def __init__(self, filepath: str, *args):
        self.message: str = f"File {filepath} not found."
        super().__init__(*args)

class NoImagesToCreateGifError(Exception):
    """0 images were provided."""
    def __init__(self, *args):
        self.message: str = "0 images were provided."
        super().__init__(*args)