build_core:
	cd ./core/Hub && gomobile bind -target android -androidapi 31 -o ../../android/app/libs/hub.aar github.com/Mihalic2040/Hub

init:
	cd ./core/Hub && gomobile init
