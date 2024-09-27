"""View module."""

# -*- coding: utf-8 -*-

################################################################################
## Form generated from reading UI file 'ui_main.ui'
##
## Created by: Qt User Interface Compiler version 6.7.2
##
## WARNING! All changes made in this file will be lost when recompiling UI file!
################################################################################

from PySide6.QtCore import QCoreApplication, QMetaObject, QRect, QSize
from PySide6.QtGui import QAction, QIcon
from PySide6.QtWidgets import (
    QFrame,
    QHBoxLayout,
    QLineEdit,
    QListWidget,
    QMenuBar,
    QPushButton,
    QSizePolicy,
    QSpacerItem,
    QVBoxLayout,
    QWidget,
)


class View(object):
    """
    View class.
    
    Sets the ui and connects it to backend.
    """
    def setupUi(self, MainWindow):
        if not MainWindow.objectName():
            MainWindow.setObjectName("MainWindow")
        MainWindow.resize(500, 800)
        MainWindow.setMinimumSize(QSize(200, 300))
        MainWindow.setMouseTracking(False)
        MainWindow.setAutoFillBackground(False)
        MainWindow.setStyleSheet("color: #d1d1cc;")
        self.choose_file_action = QAction(MainWindow)
        self.choose_file_action.setObjectName("choose_file_action")
        self.save_action = QAction(MainWindow)
        self.save_action.setObjectName("save_action")
        self.actionTrsfewfwefwe = QAction(MainWindow)
        self.actionTrsfewfwefwe.setObjectName("actionTrsfewfwefwe")
        self.actionerfgwergewrg = QAction(MainWindow)
        self.actionerfgwergewrg.setObjectName("actionerfgwergewrg")
        self.actionerfgwergewrg.setEnabled(True)
        self.actionwegfwegwreg = QAction(MainWindow)
        self.actionwegfwegwreg.setObjectName("actionwegfwegwreg")
        self.negative_action = QAction(MainWindow)
        self.negative_action.setObjectName("negative_action")
        self.negative_action.setCheckable(True)
        self.red_channel_action = QAction(MainWindow)
        self.red_channel_action.setObjectName("red_channel_action")
        self.red_channel_action.setCheckable(True)
        self.red_channel_action.setChecked(True)
        self.green_channel_action = QAction(MainWindow)
        self.green_channel_action.setObjectName("green_channel_action")
        self.green_channel_action.setCheckable(True)
        self.green_channel_action.setChecked(True)
        self.blue_channel_action = QAction(MainWindow)
        self.blue_channel_action.setObjectName("blue_channel_action")
        self.blue_channel_action.setCheckable(True)
        self.blue_channel_action.setChecked(True)
        self.action_5 = QAction(MainWindow)
        self.action_5.setObjectName("action_5")
        self.action_6 = QAction(MainWindow)
        self.action_6.setObjectName("action_6")
        self.vebcam_capture_action = QAction(MainWindow)
        self.vebcam_capture_action.setObjectName("vebcam_capture_action")
        self.about_action = QAction(MainWindow)
        self.about_action.setObjectName("about_action")
        icon = QIcon()
        icon.addFile(
            "../../../\u041e\u0437\u041f\u0420/practika1/pythonProject/gui/img/github_logo.png",
            QSize(),
            QIcon.Mode.Normal,
            QIcon.State.Off,
        )
        self.about_action.setIcon(icon)
        self.centralwidget = QWidget(MainWindow)
        self.centralwidget.setObjectName("centralwidget")
        self.centralwidget.setStyleSheet(
            "border: 1px solid #4d4e51;\n"
            "background-color: #1e1f22;\n"
            "color: #d1d1cc;"
        )
        self.verticalLayout = QVBoxLayout(self.centralwidget)
        self.verticalLayout.setSpacing(0)
        self.verticalLayout.setObjectName("verticalLayout")
        self.verticalLayout.setContentsMargins(0, 0, 0, 0)
        self.inner_frame = QFrame(self.centralwidget)
        self.inner_frame.setObjectName("inner_frame")
        self.inner_frame.setEnabled(True)
        self.inner_frame.setStyleSheet("QPushButton:hover{\n" "color: white;\n" "}")
        self.inner_frame.setFrameShape(QFrame.Shape.NoFrame)
        self.verticalLayout_2 = QVBoxLayout(self.inner_frame)
        self.verticalLayout_2.setSpacing(0)
        self.verticalLayout_2.setObjectName("verticalLayout_2")
        self.verticalLayout_2.setContentsMargins(0, 0, 0, 0)
        self.topbar = QWidget(self.inner_frame)
        self.topbar.setObjectName("topbar")
        self.topbar.setMinimumSize(QSize(0, 20))
        self.topbar.setStyleSheet("border: none")
        self.horizontalLayout_3 = QHBoxLayout(self.topbar)
        self.horizontalLayout_3.setSpacing(5)
        self.horizontalLayout_3.setObjectName("horizontalLayout_3")
        self.horizontalLayout_3.setContentsMargins(3, 0, 0, 0)
        self.previous_btn = QPushButton(self.topbar)
        self.previous_btn.setObjectName("previous_btn")

        self.horizontalLayout_3.addWidget(self.previous_btn)

        self.play_btn = QPushButton(self.topbar)
        self.play_btn.setObjectName("play_btn")

        self.horizontalLayout_3.addWidget(self.play_btn)

        self.stop_btn = QPushButton(self.topbar)
        self.stop_btn.setObjectName("stop_btn")

        self.horizontalLayout_3.addWidget(self.stop_btn)

        self.next_btn = QPushButton(self.topbar)
        self.next_btn.setObjectName("next_btn")

        self.horizontalLayout_3.addWidget(self.next_btn)

        self.horizontalSpacer = QSpacerItem(
            40, 20, QSizePolicy.Policy.Expanding, QSizePolicy.Policy.Minimum
        )

        self.horizontalLayout_3.addItem(self.horizontalSpacer)

        self.verticalLayout_2.addWidget(self.topbar)

        self.main_space_widget = QWidget(self.inner_frame)
        self.main_space_widget.setObjectName("main_space_widget")
        self.horizontalLayout = QHBoxLayout(self.main_space_widget)
        self.horizontalLayout.setObjectName("horizontalLayout")
        self.tracklist_widget = QWidget(self.main_space_widget)
        self.tracklist_widget.setObjectName("tracklist_widget")
        self.verticalLayout_4 = QVBoxLayout(self.tracklist_widget)
        self.verticalLayout_4.setObjectName("verticalLayout_4")
        self.tracklist_list = QListWidget(self.tracklist_widget)
        self.tracklist_list.setObjectName("tracklist")

        self.verticalLayout_4.addWidget(self.tracklist_list)

        self.add_button = QPushButton(self.tracklist_widget)
        self.add_button.setObjectName("add_button")

        self.verticalLayout_4.addWidget(self.add_button)

        self.remove_button = QPushButton(self.tracklist_widget)
        self.remove_button.setObjectName("remove_button")

        self.verticalLayout_4.addWidget(self.remove_button)

        self.change_button = QPushButton(self.tracklist_widget)
        self.change_button.setObjectName("change_button")

        self.verticalLayout_4.addWidget(self.change_button)

        self.horizontalLayout.addWidget(self.tracklist_widget)

        self.playlist_widget = QWidget(self.main_space_widget)
        self.playlist_widget.setObjectName("playlist_widget")
        self.verticalLayout_5 = QVBoxLayout(self.playlist_widget)
        self.verticalLayout_5.setObjectName("verticalLayout_5")
        self.playlist_list = QListWidget(self.playlist_widget)
        self.playlist_list.setObjectName("playlist_list")

        self.verticalLayout_5.addWidget(self.playlist_list)

        self.playlist_name_input = QLineEdit(self.playlist_widget)
        self.playlist_name_input.setObjectName("playlist_name_input")

        self.verticalLayout_5.addWidget(self.playlist_name_input)

        self.create_playlist_button = QPushButton(self.playlist_widget)
        self.create_playlist_button.setObjectName("create_playlist_button")

        self.verticalLayout_5.addWidget(self.create_playlist_button)

        self.delete_playlist_button = QPushButton(self.playlist_widget)
        self.delete_playlist_button.setObjectName("delete_playlist_button")

        self.verticalLayout_5.addWidget(self.delete_playlist_button)

        self.select_playlist_button = QPushButton(self.playlist_widget)
        self.select_playlist_button.setObjectName("select_playlist_button")

        self.verticalLayout_5.addWidget(self.select_playlist_button)

        self.horizontalLayout.addWidget(self.playlist_widget)

        self.verticalLayout_2.addWidget(self.main_space_widget)

        self.footer_widget = QWidget(self.inner_frame)
        self.footer_widget.setObjectName("footer_widget")
        self.footer_widget.setMinimumSize(QSize(25, 25))
        self.footer_widget.setMaximumSize(QSize(16777215, 25))
        self.footer_widget.setStyleSheet("QWidget {\n" "padding: 0;\n" "}")
        self.horizontalLayout_2 = QHBoxLayout(self.footer_widget)
        self.horizontalLayout_2.setSpacing(0)
        self.horizontalLayout_2.setObjectName("horizontalLayout_2")
        self.horizontalLayout_2.setContentsMargins(0, 0, 0, 0)
        self.lineEdit = QLineEdit(self.footer_widget)
        self.lineEdit.setObjectName("lineEdit")

        self.horizontalLayout_2.addWidget(self.lineEdit)

        self.verticalLayout_2.addWidget(self.footer_widget)

        self.verticalLayout.addWidget(self.inner_frame)

        MainWindow.setCentralWidget(self.centralwidget)
        self.menuBar = QMenuBar(MainWindow)
        self.menuBar.setObjectName("menuBar")
        self.menuBar.setEnabled(True)
        self.menuBar.setGeometry(QRect(0, 0, 836, 22))
        self.menuBar.setStyleSheet(
            "QMenuBar{\n"
            "background-color:#2b2d30;\n"
            "color: #d1d1cc;\n"
            "}\n"
            "QMenuBar::item:selected{\n"
            "color: #e2e2dd;\n"
            "background: #4d4e51;\n"
            "}\n"
            "QMenu{\n"
            "background-color:#2b2d30;\n"
            "color: #d1d1cc;\n"
            "border: 1px solid #18191c;\n"
            "}\n"
            "QMenu::item{\n"
            "background-color:#2b2d30;\n"
            "color: #d1d1cc;\n"
            "margin: 0;\n"
            "padding: 2px 8px 2px 4px;\n"
            "}\n"
            "QMenu::item:selected{\n"
            "color: #e2e2dd;\n"
            "background: #4d4e51;\n"
            "}\n"
            "\n"
            "\n"
            "\n"
            "\n"
            "\n"
            "\n"
            "\n"
            "r"
        )
        self.previous_menu_action = QAction(self.menuBar)
        self.previous_menu_action.setObjectName("previous_menu_action")
        self.play_menu_action = QAction(self.menuBar)
        self.play_menu_action.setObjectName("play_menu_action")
        self.next_menu_action = QAction(self.menuBar)
        self.next_menu_action.setObjectName("next_menu_action")
        self.next_menu_action.setEnabled(True)
        self.stop_menu_action = QAction(self.menuBar)
        self.stop_menu_action.setObjectName("stop_menu_action")
        MainWindow.setMenuBar(self.menuBar)

        self.menuBar.addAction(self.previous_menu_action)
        self.menuBar.addAction(self.play_menu_action)
        self.menuBar.addAction(self.stop_menu_action)
        self.menuBar.addAction(self.next_menu_action)

        self.retranslateUi(MainWindow)

        QMetaObject.connectSlotsByName(MainWindow)

    # setupUi

    def retranslateUi(self, MainWindow):
        MainWindow.setWindowTitle(
            QCoreApplication.translate("MainWindow", "MainWindow", None)
        )
        self.choose_file_action.setText(
            QCoreApplication.translate(
                "MainWindow",
                "\u0412\u044b\u0431\u0440\u0430\u0442\u044c \u0444\u0430\u0439\u043b...",
                None,
            )
        )
        self.save_action.setText(
            QCoreApplication.translate(
                "MainWindow",
                "\u0421\u043e\u0445\u0440\u0430\u043d\u0438\u0442\u044c",
                None,
            )
        )
        self.actionTrsfewfwefwe.setText(
            QCoreApplication.translate("MainWindow", "Trsfewfwefwe", None)
        )
        self.previous_btn.setText(QCoreApplication.translate("MainWindow", "<<", None))
        self.play_btn.setText(QCoreApplication.translate("MainWindow", "play", None))
        self.stop_btn.setText(QCoreApplication.translate("MainWindow", "stop", None))
        self.next_btn.setText(QCoreApplication.translate("MainWindow", ">>", None))
        self.actionerfgwergewrg.setText(
            QCoreApplication.translate("MainWindow", "erfgwergewrg", None)
        )
        self.actionwegfwegwreg.setText(
            QCoreApplication.translate("MainWindow", "wegfwegwreg", None)
        )
        self.negative_action.setText(
            QCoreApplication.translate(
                "MainWindow", "\u041d\u0435\u0433\u0430\u0442\u0438\u0432", None
            )
        )
        self.red_channel_action.setText(
            QCoreApplication.translate(
                "MainWindow", "\u041a\u0440\u0430\u0441\u043d\u044b\u0439", None
            )
        )
        self.green_channel_action.setText(
            QCoreApplication.translate(
                "MainWindow", "\u0417\u0435\u043b\u0451\u043d\u044b\u0439", None
            )
        )
        self.blue_channel_action.setText(
            QCoreApplication.translate(
                "MainWindow", "\u0421\u0438\u043d\u0438\u0439", None
            )
        )
        self.action_5.setText(
            QCoreApplication.translate(
                "MainWindow",
                "\u043f\u043e\u0432\u0435\u0440\u043d\u0443\u0442\u044c \u0432\u043b\u0435\u0432\u043e (5 \u0433\u0440\u0430\u0434\u0443\u0441\u043e\u0432)",
                None,
            )
        )
        self.action_6.setText(
            QCoreApplication.translate(
                "MainWindow",
                "\u043f\u043e\u0432\u0435\u0440\u043d\u0443\u0442\u044c \u0432\u043f\u0440\u0430\u0432\u043e (5 \u0433\u0440\u0430\u0434\u0443\u0441\u043e\u0432)",
                None,
            )
        )
        self.vebcam_capture_action.setText(
            QCoreApplication.translate(
                "MainWindow",
                "\u0421\u043d\u044f\u0442\u044c \u0438\u0437\u043e\u0431\u0440\u0430\u0436\u0435\u043d\u0438\u0435 \u0441 \u043a\u0430\u043c\u0435\u0440\u044b",
                None,
            )
        )
        self.about_action.setText(
            QCoreApplication.translate(
                "MainWindow",
                "\u041e \u043f\u0440\u043e\u0433\u0440\u0430\u043c\u043c\u0435...",
                None,
            )
        )
        self.add_button.setText(
            QCoreApplication.translate(
                "MainWindow",
                "\u0414\u043e\u0431\u0430\u0432\u0438\u0442\u044c \u0442\u0440\u0435\u043a",
                None,
            )
        )
        self.remove_button.setText(
            QCoreApplication.translate(
                "MainWindow",
                "\u0423\u0434\u0430\u043b\u0438\u0442\u044c \u0442\u0440\u0435\u043a",
                None,
            )
        )
        self.change_button.setText(
            QCoreApplication.translate(
                "MainWindow",
                "\u041f\u043e\u043c\u0435\u043d\u044f\u0442\u044c \u043f\u043e\u0437\u0438\u0446\u0438\u044e",
                None,
            )
        )
        self.create_playlist_button.setText(
            QCoreApplication.translate(
                "MainWindow",
                "\u0421\u043e\u0437\u0434\u0430\u0442\u044c \u043f\u043b\u0435\u0439\u043b\u0438\u0441\u0442",
                None,
            )
        )
        self.delete_playlist_button.setText(
            QCoreApplication.translate(
                "MainWindow",
                "\u0423\u0434\u0430\u043b\u0438\u0442\u044c \u043f\u043b\u0435\u0439\u043b\u0438\u0441\u0442",
                None,
            )
        )
        self.select_playlist_button.setText(
            QCoreApplication.translate(
                "MainWindow",
                "\u0412\u044b\u0431\u0440\u0430\u0442\u044c \u043f\u043b\u0435\u0439\u043b\u0438\u0441\u0442",
                None,
            )
        )
        self.lineEdit.setText(
            QCoreApplication.translate(
                "MainWindow",
                "https://github.com/egorgur/aads/tree/main/practice/practice_1.2",
                None,
            )
        )
        # self.previous_menu_action.(
        #     QCoreApplication.translate("MainWindow", "<<", None)
        # )
        # self.play_menu_action.setTitle(
        #     QCoreApplication.translate("MainWindow", "play", None)
        # )
        # self.next_menu_action.setTitle(
        #     QCoreApplication.translate("MainWindow", ">>", None)
        # )
        # self.stop_menu_action.setTitle(
        #     QCoreApplication.translate("MainWindow", "stop", None)
        # )

    # retranslateUi

    def set_connections(self):
        """Bind the callback functions to ui."""

        self.add_button.clicked.connect(self.add_song)

        self.remove_button.clicked.connect(self.remove_song)

        self.play_btn.clicked.connect(self.choose_song)

        self.stop_btn.clicked.connect(self.stop_track)

        self.previous_btn.clicked.connect(self.play_previous)

        self.next_btn.clicked.connect(self.play_next)

        self.change_button.clicked.connect(self.change_position_dialogue)


        self.create_playlist_button.clicked.connect(self.create_playlist)

        self.delete_playlist_button.clicked.connect(self.delete_playlist)

        self.select_playlist_button.clicked.connect(self.select_playlist)

        