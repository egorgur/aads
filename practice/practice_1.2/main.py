"""Модуль интерфейса"""

import os
import sys

import pygame
from music_track import MusicTrack
from playlist import Playlist
from PySide6 import QtCore
from PySide6.QtWidgets import (
    QApplication,
    QFileDialog,
    QInputDialog,
    QMainWindow,
    QMessageBox,
)
from view import View

# class PlaylistManager(QDialog):
#     """PlayList manager UI."""

#     add_playlist = QtCore.Signal(str)
#     remove_playlist = QtCore.Signal(str)
#     choose_playlist = QtCore.Signal(str)

#     def __init__(self, playlist_objects: list[Playlist]):
#         super().__init__()
#         self.playlist_objects = playlist_objects
#         self.setWindowTitle("Playlist Manager")
#         self.setGeometry(100, 100, 300, 300)

#         self.playlist_list = QListWidget()
#         self.playlist_name_input = QLineEdit()
#         self.playlist_name_input.setPlaceholderText("Введите название плейлиста")

#         self.create_button = QPushButton("Создать плейлист")
#         self.delete_button = QPushButton("Удалить плейлист")
#         self.select_button = QPushButton("Выбрать плейлист")

#         self.create_button.clicked.connect(self.create_playlist)
#         self.delete_button.clicked.connect(self.delete_playlist)
#         self.select_button.clicked.connect(self.select_playlist)

#         layout = QVBoxLayout()
#         layout.addWidget(self.playlist_list)
#         layout.addWidget(self.playlist_name_input)
#         layout.addWidget(self.create_button)
#         layout.addWidget(self.delete_button)
#         layout.addWidget(self.select_button)

#         self.setLayout(layout)

#         if len(self.playlist_objects) != 0:
#             for playlist_object in self.playlist_objects:
#                 self.playlist_list.addItem(
#                     f"{playlist_object["name"]}, tracks: {len(playlist_object["playlist"])}"
#                 )

#     def create_playlist(self):
#         """Create playlist."""
#         name = self.playlist_name_input.text()
#         if name:
#             if "," in name:
#                 QMessageBox.warning(
#                     self, "Ошибка", "Название плейлиста не должно содержать запятую."
#                 )
#                 return

#             display_name = f"{name}, композиций: {0}"

#             if get_playlist_object_by_name(display_name, self.playlist_objects):
#                 QMessageBox.warning(
#                     self, "Ошибка", "Плейлист с таким названием уже существует"
#                 )
#                 return

#             self.playlist_list.addItem(display_name)
#             self.playlist_name_input.clear()
#             self.add_playlist.emit(name)
#         else:
#             QMessageBox.warning(self, "Ошибка", "Вы не ввели название плейлиста")

#     def delete_playlist(self):
#         """Delete playlist."""
#         selected_items = self.playlist_list.selectedItems()
#         if not selected_items:
#             QMessageBox.warning(self, "Ошибка", "Плейлист не выбран")
#             return

#         for item in selected_items:
#             self.playlist_list.takeItem(self.playlist_list.row(item))
#             self.remove_playlist.emit(item.text())

#     def select_playlist(self):
#         """Select playlist."""
#         selected_items = self.playlist_list.selectedItems()
#         if not selected_items:
#             QMessageBox.warning(self, "Ошибка", "Плейлист не выбран")
#             return

#         selected_playlist = selected_items[0].text()
#         self.choose_playlist.emit(selected_playlist)
#         self.close()


class MusicPlayer(QMainWindow, View):
    """Music Player Window."""

    add_playlist = QtCore.Signal(str)
    remove_playlist = QtCore.Signal(str)
    choose_playlist = QtCore.Signal(str)

    def __init__(self):
        super().__init__()
        self.setupUi(self)
        self.set_connections()
        self.current_playlist: Playlist = None
        self.playlist_objects = []
        self.playlist_manager = None

        pygame.mixer.init()

        self.selected_song_index = None

        self.timer = QtCore.QTimer(self)
        self.timer.setInterval(1000)
        self.timer.timeout.connect(self.next_track_if_this_ended)

        self.choose_playlist.connect(self.on_select_playlist)
        self.remove_playlist.connect(self.on_remove_playlist)
        self.add_playlist.connect(self.on_create_playlist)
    
    def create_playlist(self):
        """Create playlist."""
        name = self.playlist_name_input.text()
        if name:
            if "," in name:
                QMessageBox.warning(
                    self, "Ошибка", "Название плейлиста не должно содержать запятую."
                )
                return

            display_name = f"{name}"

            if get_playlist_object_by_name(display_name, self.playlist_objects):
                QMessageBox.warning(
                    self, "Ошибка", "Плейлист с таким названием уже существует"
                )
                return

            self.playlist_list.addItem(display_name)
            self.playlist_name_input.clear()
            self.add_playlist.emit(name)
        else:
            QMessageBox.warning(self, "Ошибка", "Вы не ввели название плейлиста")

    def delete_playlist(self):
        """Delete playlist."""
        selected_items = self.playlist_list.selectedItems()
        if not selected_items:
            QMessageBox.warning(self, "Ошибка", "Плейлист не выбран")
            return

        for item in selected_items:
            self.playlist_list.takeItem(self.playlist_list.row(item))
            self.remove_playlist.emit(item.text())

    def select_playlist(self):
        """Select playlist."""
        selected_items = self.playlist_list.selectedItems()
        if not selected_items:
            QMessageBox.warning(self, "Ошибка", "Плейлист не выбран")
            return

        selected_playlist = selected_items[0].text()
        self.choose_playlist.emit(selected_playlist)

    def on_select_playlist(self, name):
        """Playlist select event handler."""
        print("efwefewfwgwegewgewgwwegwegwegweg")
        self.current_playlist = get_playlist_object_by_name(
            name, self.playlist_objects
        )["playlist"]
        print(self.current_playlist)
        self.tracklist_list.clear()
        for node in self.current_playlist:
            self.tracklist_list.addItem(os.path.basename(node.data.path))

    def on_remove_playlist(self, name):
        """Delete playlist event handler."""
        print(self.playlist_objects)
        self.playlist_objects.remove(
            get_playlist_object_by_name(name, self.playlist_objects)
        )

    def on_create_playlist(self, name):
        """Create playlist event handler."""
        print(self.playlist_objects)
        self.playlist_objects.append({"name": name, "playlist": Playlist()})

    def stop_track(self):
        """Stop current track."""
        self.current_playlist.current = None
        pygame.mixer.music.unload()
        pygame.mixer.music.stop()

    def add_song(self):
        """Add a song."""
        if self.current_playlist is None:
            QMessageBox.warning(self, "Ошибка", "Сначала выберите плейлист")
            return

        options = QFileDialog.Options()
        files, _ = QFileDialog.getOpenFileNames(
            self,
            "Выберите музыкальные файлы",
            "",
            "Audio Files (*.mp3 *.wav)",
            options=options,
        )
        if files:
            for file in files:
                self.current_playlist.append(MusicTrack(file))
                self.tracklist_list.addItem(os.path.basename(file))

    def remove_song(self):
        """Remove a song."""
        selected_items = self.tracklist_list.selectedItems()
        if not selected_items:
            QMessageBox.warning(self, "Ошибка", "Выберите композицию для удаления.")
            return
        for item in selected_items:
            row = self.tracklist_list.row(item)
            if self.current_playlist.current == self.current_playlist[row]:
                self.current_playlist.current = None
                pygame.mixer.music.unload()
                pygame.mixer.music.stop()

            self.current_playlist.remove(self.current_playlist[row])
            self.tracklist_list.takeItem(row)

    def find_song_by_id(self, sid):
        """Find song by id."""
        for idx, track in enumerate(self.current_playlist):
            if idx == sid:
                return track

        return None

    def play_song_by_id(self, sid):
        """Play song by id."""
        track = self.find_song_by_id(sid)
        self.current_playlist.play_all(track)
        self.timer.start()

    def choose_song(self):
        """Play song."""
        current_row = self.tracklist_list.currentRow()
        if current_row < 0:
            QMessageBox.warning(
                self, "Ошибка", "Выберите композицию для воспроизведения."
            )
            return None

        if self.current_playlist[current_row] != self.current_playlist.current:
            return self.play_song_by_id(current_row)

    def play_previous(self):
        """Play previous song."""
        if self.current_playlist is None:
            QMessageBox.warning(self, "Ошибка", "Сначала выберите плейлист")
            return

        if not pygame.mixer.music.get_busy():
            QMessageBox.warning(self, "Ошибка", "Сейчас не проигрывается трек")
            return

        self.current_playlist.previous_track()

    def play_next(self):
        """Play next track."""
        if self.current_playlist is None:
            QMessageBox.warning(self, "Ошибка", "Сначала выберите плейлист")
            return

        if not pygame.mixer.music.get_busy():
            QMessageBox.warning(self, "Ошибка", "Сейчас не проигрывается трек")
            return

        self.current_playlist.next_track()

    def next_track_if_this_ended(self):
        """Play next track on current end."""
        if pygame.mixer.music.get_busy():
            return

        self.current_playlist.next_track()

    def change_position_dialogue(self) -> int:
        """Change track position."""
        current_row = self.tracklist_list.currentRow()
        if current_row < 0:
            QMessageBox.warning(
                self, "Ошибка", "Выберите композицию для смены позиции."
            )
            return

        if self.current_playlist is None:
            QMessageBox.warning(self, "Ошибка", "Сначала выберите плейлист")
            return

        if len(self.current_playlist) < 2:
            QMessageBox.warning(
                self,
                "Ошибка",
                "В плейлисте должно быть минимум 2 композиции для перестановки",
            )
            return

        response = QInputDialog.getInt(
            self,
            "Место композиции",
            f"Введите число от 1 до {len(self.current_playlist) - 1}",
            minValue=1,
            maxValue=len(self.current_playlist) - 1,
        )

        if response[1] is False:
            return

        number = response[0]

        if current_row == number:
            return

        self.change_position(number, current_row)

    def change_position(self, position, current_row):
        """Смена позиции песни"""
        selected_track = self.find_song_by_id(current_row)
        prev_track = self.find_song_by_id(position - 1)

        if current_row < position:
            prev_track = self.find_song_by_id(position)

        self.remove_song()
        self.current_playlist.insert(prev_track.data, selected_track.data)
        self.tracklist_list.clear()
        for node in self.current_playlist:
            self.tracklist_list.addItem(os.path.basename(node.data.path))


def get_playlist_object_by_name(name: str, playlist_objects: list):
    """Получение плейлиста-объекта по названию"""
    for playlist_object in playlist_objects:
        if playlist_object["name"] == name:
            return playlist_object

    return None


if __name__ == "__main__":
    app = QApplication(sys.argv)
    player = MusicPlayer()
    player.show()
    sys.exit(app.exec_())
