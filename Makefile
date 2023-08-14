push:
	git branch | grep "*" | sed -e "s/* //g" | xargs git push origin
