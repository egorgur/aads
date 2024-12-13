"""Fast api service."""

from fastapi import FastAPI
from views import router as api_routes

app = FastAPI()
app.include_router(api_routes)
"""Fast API application."""
