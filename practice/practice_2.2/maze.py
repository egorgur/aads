"""
Maze class module.
"""

import copy
import random
from logging import Logger

import exceptions
from PIL import Image, ImageDraw


def shortest_distance(start_pos: tuple[int, int], end_pos: tuple[int, int]) -> int:
    """
    Calculate the shortest distance between two positions.

    Args:
        start_pos (Tuple[int, int]): The starting position as a (row, column) tuple.
        end_pos (Tuple[int, int]): The ending position as a (row, column) tuple.

    Returns:
        int: The shortest distance between the positions.
    """
    return (abs(start_pos[0] - end_pos[0]) + abs(start_pos[1] - end_pos[1])) // 2


def find_element_in_matrix(matrix: list[list[int]], target: int) -> list[int] | None:
    """
    Find the target value in a 2D matrix.

    Args:
        matrix (List[List[int]]): The 2D matrix to search.
        target (int): The target value to find.

    Returns:
        Union[List[int], None]: The position of the target value as a (row, column) tuple or None if not found.
    """
    for i in range(len(matrix)):
        for j in range(len(matrix[i])):
            if matrix[i][j] == target:
                return [i, j]
    return None


class Maze:
    """
    Class that represents a maze.
    """

    def __init__(self, logger: Logger, rows: int = 1, cols: int = 1) -> None:
        """Initialize a Maze instance."""
        self.rows: int = rows * 2 + 1
        self.cols: int = cols * 2 + 1
        # rows and cols are doubled and incremented once to create walls
        self.maze: list[list[int]] | None = None

        self.index: int = 2
        # Special parameter to determine the wall breaking/creation

        self.generation_steps: list[list[list[int]]] = []
        # List of generation steps

        self.path: list[list[int]] = None
        # list of path cells

        self.solving_steps: list[list[list[int]]] = []
        # List of solving steps

    def generate(self) -> None:
        """Generate a maze."""

        self.maze = [[0] * self.cols for _ in range(self.rows)]

        # == COPY maze ==
        self.generation_steps.append(copy.deepcopy(self.maze))

        # Borderline
        self.maze = [
            [
                1 if row in (0, self.rows - 1) or col in (0, self.cols - 1) else 0
                for col in range(self.cols)
            ]
            for row in range(self.rows)
        ]

        # == COPY maze ==
        self.generation_steps.append(copy.deepcopy(self.maze))

        # Set grid-like walls to build upon
        for row in range(2, self.rows, 2):
            for col in range(0, self.cols, 2):
                self.maze[row][col] = 1

        # == COPY maze ==
        self.generation_steps.append(copy.deepcopy(self.maze))

        # first row indicies
        for col in range(1, self.cols, 2):
            self.maze[1][col] = self.index
            self.index += 1

        for row in range(1, self.rows, 2):
            # right walls
            for col in range(1, self.cols - 2, 2):
                if random.choice((True, False)):
                    self.maze[row][col + 1] = 1
                else:
                    if self.maze[row][col] == self.maze[row][col + 2]:
                        self.maze[row][col + 1] = 1
                    else:
                        temp = copy.copy(self.maze[row][col + 2])

                        for temp_col in range(1, self.cols, 2):
                            if self.maze[row][temp_col] == temp:
                                self.maze[row][temp_col] = self.maze[row][col]

            # bottom walls
            for col in range(1, self.cols, 2):
                if row != self.rows - 2:
                    self.maze[row + 2][col] = self.maze[row][col]
                    if random.choice((True, False)):
                        # Place wall under the cell if exist an exit with the same index
                        count = 0
                        temp = copy.copy(self.maze[row][col])
                        for temp_col in range(1, self.cols, 2):
                            if (
                                self.maze[row][temp_col] == temp
                                and self.maze[row + 1][temp_col] == 0
                            ):
                                count += 1
                        if count > 1:
                            self.maze[row + 1][col] = 1
                            self.maze[row + 2][col] = self.index
                            self.index += 1

            # == COPY maze == on each row
            self.generation_steps.append(copy.deepcopy(self.maze))

        # Last line
        for col in range(1, self.cols - 2, 2):
            if self.maze[self.rows - 2][col] != self.maze[self.rows - 2][col + 2]:
                self.maze[self.rows - 2][col + 1] = 0
                temp = copy.copy(self.maze[self.rows - 2][col + 2])
                for temp_col in range(1, self.cols, 2):
                    if self.maze[self.rows - 2][temp_col] == temp:
                        self.maze[self.rows - 2][temp_col] = self.maze[self.rows - 2][
                            temp_col
                        ]
        # == COPY maze ==
        self.generation_steps.append(copy.deepcopy(self.maze))
        # delete trash
        for row in range(1, self.rows, 2):
            for col in range(1, self.cols, 2):
                self.maze[row][col] = 0
        self.index = 2

    def solve(self, start: tuple[int, int], end: tuple[int, int]) -> None:
        """Solve recursively of the maze from a given start position to an end position."""
        if not (all(x % 2 != 0 for x in start) and all(x % 2 != 0 for x in end)):
            return
        if start == end:
            return
        max_value = max(self.rows, self.cols)
        if not (
            1 <= start[0] <= max_value
            and 1 <= start[1] <= max_value
            and 1 <= end[0] <= max_value
            and 1 <= end[1] <= max_value
        ):
            return

        def a_way_out(
            maze: list[list[list[int]]],
            start_pos: tuple[int, int],
            end_pos: tuple[int, int],
        ) -> None:
            """
            Recursive function to find a way out in the maze using backtracking.

            Args:
                maze (List[List[List[int]]): List representing the maze layout.
                start_pos (Tuple[int, int]): The current position as a (row, column) tuple.
                end_pos (Tuple[int, int]): The destination position as a (row, column) tuple.

            Returns:
                None
            """
            maze[start_pos[0]][start_pos[1]][1] = 1
            self.solving_steps.append(maze[start_pos[0]][start_pos[1]][2])
            ways = []

            def try_move(direction: tuple[int, int], distance: int) -> None:
                """
                Attempt to move in a specific direction from the current position and update the maze state.

                Args:
                    maze (List[List[List[int]]): List representing the maze layout.
                    direction (Tuple[int, int]): The movement direction.
                    distance (int): The distance to move in that direction.

                Returns:
                    None
                """
                new_pos_wall = (
                    start_pos[0] + direction[0],
                    start_pos[1] + direction[1],
                )
                new_pos = (
                    start_pos[0] + direction[0] * distance,
                    start_pos[1] + direction[1] * distance,
                )
                if (
                    maze[new_pos_wall[0]][new_pos_wall[1]] != 1
                    and maze[new_pos[0]][new_pos[1]][1] != 1
                ):
                    maze[new_pos[0]][new_pos[1]] = [
                        shortest_distance(new_pos, end_pos),
                        0,
                        maze[start_pos[0]][start_pos[1]][2]
                        + [[new_pos[0], new_pos[1]]],
                    ]
                    ways.append(maze[new_pos[0]][new_pos[1]])

            try_move((0, 1), 2)
            try_move((-1, 0), 2)
            try_move((1, 0), 2)
            try_move((0, -1), 2)
            shortest_ways = list(filter(lambda x: not x[1], ways))
            shortest_ways.sort(key=lambda x: x[0])
            if any(sublist[:2] == [0, 0] for sublist in shortest_ways):
                return
            if shortest_ways:
                new_start = find_element_in_matrix(maze, shortest_ways[0])
                a_way_out(maze, new_start, end_pos)
            else:
                new_start = [1, 1]
                for i in range(1, self.rows, 2):
                    for j in range(1, self.cols, 2):
                        if maze[i][j][0] != 0 and maze[i][j][1] != 1:
                            if maze[i][j][0] < maze[new_start[0]][new_start[1]][0]:
                                new_start = [i, j]
                a_way_out(maze, new_start, end_pos)

        solving_maze = copy.deepcopy(self.maze)
        for i in range(1, self.rows, 2):
            for j in range(1, self.cols, 2):
                solving_maze[i][j] = [0, 0, 0]
        solving_maze[start[0]][start[1]] = [
            shortest_distance(start, end),
            0,
            [list(start)],
        ]
        self.solving_steps.append(solving_maze[start[0]][start[1]][2])
        a_way_out(solving_maze, start, end)
        self.solving_steps.append(solving_maze[end[0]][end[1]][2])
        self.solving_steps.append(solving_maze[end[0]][end[1]][2])
        self.solving_steps.append(solving_maze[end[0]][end[1]][2])
        self.path = solving_maze[end[0]][end[1]][2]

    def import_maze_from_file(self, filename: str) -> None:
        """
        Import a maze from a text file.

        Args:
            filename (str): The name of the text file containing the maze.

        Returns:
            None
        """
        try:
            with open(filename, "r", encoding="utf-8") as file:
                maze_data = [list(map(int, line.strip())) for line in file.readlines()]
                self.maze = maze_data
                self.rows = len(maze_data)
                self.cols = len(maze_data[0])

        except FileNotFoundError:
            raise exceptions.MazeFileNotFoundError(filepath=filename)

    def import_maze_from_image(self, filename: str) -> None:
        """
        Import a maze from an image file.

        Args:
            filename (str): The name of the image file containing the maze.

        Returns:
            None
        """
        wall_color = (0, 0, 0)
        path_color = (255, 255, 255)
        try:
            image = Image.open(filename)
            width, height = image.size

            maze_data = []
            for y in range(0, height, 21):
                row = []
                for x in range(0, width, 21):
                    pixel = image.getpixel((x, y))
                    if pixel == wall_color:
                        row.append(1)
                    elif pixel == path_color:
                        row.append(0)
                    else:
                        raise ValueError("Unknown pixel color in the image")
                maze_data.append(row)
            self.maze = maze_data
            self.rows = len(maze_data)
            self.cols = len(maze_data[0])
        except FileNotFoundError:
            raise exceptions.MazeFileNotFoundError(filepath=filename)
    
    def export_maze_to_file(self, filename: str) -> None:
        """
        Export the maze to a text file.

        Args:
            filename (str): The name of the text file to save the maze.

        Returns:
            None
        """
        with open(filename, 'w', encoding='utf-8') as file:
            for row in self.maze:
                file.write(''.join(map(str, row)) + '\n')
    
    def create_maze_png(self, maze: list[list[int]], solve_path: list[list[int]] = None) -> Image.Image:
        """
        Create an image of a labyrinth with a solution path if there is one.

        Args:
            maze (List[List[int]]): The maze layout.
            solve_path (List[List[int]]): The solution path as a list of (row, column) tuples (default: None).

        Returns:
            Image.Image: A PIL image representing the maze.
        """
        cell_size = 20  # Adjust cell size as needed, you should change the for algorithm in import also
        wall_color = (0, 0, 0)  # Color for walls
        path_color = (255, 255, 255)  # Color for paths
        find_color = (255, 0, 0)  # Color for searching algorithm

        width = self.cols * cell_size
        height = self.rows * cell_size
        img = Image.new('RGB', (width, height), path_color)
        draw = ImageDraw.Draw(img)

        for i in range(self.rows):
            for j in range(self.cols):
                if maze[i][j] == 1:
                    draw.rectangle([(j * cell_size, i * cell_size),
                                    ((j + 1) * cell_size, (i + 1) * cell_size)], fill=wall_color)
        if solve_path:
            for position in solve_path:
                draw.rectangle([(position[1] * cell_size, position[0] * cell_size),
                                ((position[1] + 1) * cell_size, (position[0] + 1) * cell_size)], fill=find_color)

        return img

    def create_gif_maze_gen(self, filename: str, duration: int = 1000, loop: int = 0) -> None:
        """
        Create a GIF animation of the maze generation process.

        Args:
            filename (str): The name of the GIF file to save.
            duration (int): The duration of each frame (default: 1000).
            loop (int): Number of times the GIF should loop (0 for infinite loop, default: 0).

        Returns:
            None
        """
        gif = []
        for maze_state in self.generation_steps:
            image = self.create_maze_png(maze_state)
            gif.append(image)
        gif.reverse()
        if len(gif) > 0:
            gif.reverse()
            gif[0].save(
                filename,
                save_all=True,
                append_images=gif[1:],
                duration=duration, loop=loop)
        else:
            raise exceptions.NoImagesToCreateGifError

    def create_gif_maze_solve(self, filename: str, duration: int = 1000, loop: int = 0) -> None:
        """
        Create a GIF animation of the maze solving process.

        Args:
            filename (str): The name of the GIF file to save.
            duration (int): The duration  of each frame (default: 1000).
            loop (int): Number of times the GIF should loop (0 for infinite loop, default: 0).

        Returns:
            None
        """
        gif = []

        for path in self.solving_states:
            image = self.create_maze_png(self.maze, path)
            gif.append(image)
        gif.reverse()
        if len(gif) > 0:
            gif.reverse()
            gif[0].save(
                filename,
                save_all=True,
                append_images=gif[1:],
                duration=duration, loop=loop)
        else:
            raise exceptions.NoImagesToCreateGifError

    def print(self) -> None:
        """Print the maze in console."""

        for row in self.maze:
            print("\n", end="")
            for cell in row:
                # print(cell, end="")
                if cell == 1:
                    print("█", end="")
                else:
                    print(" ", end="")
        print("\n")

    def print_solved_maze(self) -> None:
        """Print the solved maze layout in console."""
        for row in range(self.rows):
            print("\n", end="")
            for col in range(self.cols):
                if [row, col] in self.path:
                    print("~",end="")               
                if self.maze[row][col] == 1:
                    print("█", end="")
                else:
                    print(" ", end="")
        print("\n") 