"""Fast API enpoints handlers."""

import exceptions

# Nu kak eshe inache~~
from cli import logger
from fastapi import APIRouter, HTTPException, status
from maze import Maze
from pydantic import BaseModel

router = APIRouter()


maze: Maze | None = None
"""Maze object"""


class GenerationParameters(BaseModel):
    """Parameters for maze generation."""

    rows: int  # Wifht of the maze
    cols: int  # Height of the maze
    file: str | None = None
    console: bool | None = None
    text: bool | None = None
    image: bool | None = None
    gif: bool | None = None


class SolvingParameters(BaseModel):
    """Parameters for maze solving."""

    start_x: int  # Start position width
    start_y: int  # Start position height
    end_x: int  # Start position width
    end_y: int  # Start position height
    file: str | None = None
    console: bool | None = None
    text: bool | None = None
    image: bool | None = None
    gif: bool | None = None


@router.get("/")
def read_root():
    """Test endpoint."""
    return {"Hello": "Mind"}


@router.post("/generate/")
def generate(generation_parameters: GenerationParameters):
    """
    Generate maze with a given parameters.
    """
    global maze  # krinzh

    if generation_parameters.rows is None or generation_parameters.cols is None:
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail=f"Error: {exceptions.UnprovidedMazeDataError().message}",
        )

    if generation_parameters.rows * generation_parameters.cols <= 9:
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail=f"Error: {exceptions.WrongMazeDimensionsError().message}",
        )

    maze = Maze(
        logger=logger,
        rows=generation_parameters.rows,
        cols=generation_parameters.cols,
    )
    # ~~Generiruem mei3~~
    maze.generate()

    if maze is None:
        # Check is somehow something brokes for some reason
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail=f"Error: {exceptions.UnprovidedMazeDataError().message}",
        )

    if generation_parameters.console:
        maze.print()

    if generation_parameters.text:
        # text
        maze.export_maze_to_file("output.txt")

    if generation_parameters.image:
        # image
        maze.create_maze_png(maze.maze).save("output.png", "PNG")

    if generation_parameters.gif:
        # gif
        if maze.generation_steps:
            maze.create_gif_maze_gen("output_generation.gif")
        # if maze.solving_steps:
        #     maze.create_gif_maze_solve("output_solving.gif", duration=300)

    return {
        "maze": maze.maze,
        "generation_steps": maze.generation_steps,
        "rows": maze.rows,
        "cols": maze.cols,
    }


@router.post("/solve/")
def solve(solving_parameters: SolvingParameters):
    """
    Solve maze with a given parameters.
    """
    global maze  # krinzh

    if solving_parameters.file:
        # Path to a maze source-file
        maze = Maze(logger=logger)
        if solving_parameters.file.endswith(".png"):
            maze.import_maze_from_image(solving_parameters.file)
        elif solving_parameters.file.endswith(".txt"):
            maze.import_maze_from_file(solving_parameters.file)
        else:
            raise HTTPException(
                status_code=status.HTTP_400_BAD_REQUEST,
                detail=f"Error: {exceptions.UnsupportedFileFormatError().message}",
            )

    if maze is None:
        # Check if no maze were given
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail=f"Error: {exceptions.UnprovidedMazeDataError().message}",
        )

    if not (
        solving_parameters.start_x
        and solving_parameters.start_y
        and solving_parameters.end_x
        and solving_parameters.end_y
    ):
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail=f"Error: {exceptions.WrongSolvingIndiciesError().message}",
        )

    start, end = (
        (solving_parameters.start_x, solving_parameters.start_y),
        (solving_parameters.end_x, solving_parameters.end_y),
    )
    maze.solve(start, end)

    return {
        "solving_steps": maze.solving_steps,
        "rows": maze.rows,
        "cols": maze.cols,
    }
