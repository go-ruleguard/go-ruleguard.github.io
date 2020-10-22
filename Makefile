.PHONY: play

play:
	rm -rf play/
	mkdir -p play
	git clone https://github.com/life4/ruleguard-playground.git
	cd ruleguard-playground && ./build.sh
	mv ruleguard-playground/frontend/* play/
	mv ruleguard-playground/public/* play/
	rm -rf ruleguard-playground
