"""
CLI module.
"""

import argparse
import logging

import exceptions
from maze import Maze

logger = logging.getLogger(__name__)
logging.basicConfig(level=logging.DEBUG, encoding="utf-8")


def main():
    """Main function."""
    parser = argparse.ArgumentParser(
        description="Maze generation and solving CLI application."
    )
    parser.add_argument(
        "--size",
        type=str,
        help="Maze dimensions. Format '<rows>,<columns>'",
    )
    parser.add_argument(
        "--solve_indicies",
        type=str,
        help="Position of the start point and the finish point. Format '<start_row>,<start_col>,<end_row>,<end_col>'",
    )
    parser.add_argument(
        "--file",
        type=str,
        help="Path to the import file (use .png for images and .txt for text)",
    )
    parser.add_argument(
        "--console",
        action="store_true",
        help="Output the maze in console",
    )
    parser.add_argument(
        "--text",
        action="store_true",
        help="Output the maze in textfile",
    )
    parser.add_argument(
        "--image",
        action="store_true",
        help="Output the maze in image",
    )
    parser.add_argument(
        "--gif",
        action="store_true",
        help="Output the gif(s) of maze",
    )

    args = parser.parse_args()
    logger.debug(str(args))
    maze = None

    if args.size:
        # Maze dimensions
        size = args.size.split(",")
        if len(size) != 2:
            raise exceptions.WrongMazeDimensionsError
        try:
            rows, cols = map(int, size)
        except Exception:
            raise exceptions.WrongMazeDimensionsError
        maze = Maze(logger=logger, rows=rows, cols=cols)
        maze.generate()

    if args.file:
        # Path to a maze source-file
        maze = Maze(logger=logger)
        if args.file.endswith(".png"):
            maze.import_maze_from_image(args.file)
        elif args.file.endswith(".txt"):
            maze.import_maze_from_file(args.file)
        else:
            raise exceptions.UnsupportedFileFormatError

    if maze is None:
        # If no generation parameters were not provided
        # or the existing maze is not given
        raise exceptions.UnprovidedMazeDataError

    if args.solve_indicies:
        indicies = args.solve_indicies.split(",")
        if len(args) != 4:
            raise exceptions.WrongSolvingIndiciesError
        start, end = list(map(int, indicies[:2])), list(map(int, indicies[2:]))
        maze.solve_maze(start, end)

    if args.console:
        maze.print()

    if args.text:
        maze.export_maze_to_file("output.txt")
    if args.image:
        maze.create_maze_png(maze.maze).save("output.png", "PNG")
    if args.gif:
        if maze.generation_steps:
            maze.create_gif_maze_gen("output_generation.gif")
        if maze.solving_steps:
            maze.create_gif_maze_solve("output_solving.gif", duration=300)


if __name__ == "__main__":
    """CLI app execution."""
    try:
        main()
    except Exception as e:
        logger.fatal(e.message)
