# 设置中文
flatpak config --set languages "zh;en"
flatpak update

# 设置暗色
flatpak override --user --env=GTK_THEME=Adwaita-dark
flatpak install flathub org.gtk.Gtk3theme.Adwaita-dark
sudo flatpak override --filesystem=xdg-config/gtk-3.0:ro
sudo flatpak override --filesystem=xdg-config/gtk-4.0:ro
