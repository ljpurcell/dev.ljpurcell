document.addEventListener("DOMContentLoaded", function() {
	const clipboard = document.getElementById("email-clipboard");
	if (clipboard) {
		clipboard.addEventListener("click", copyEmailOnAboutPage)
	}
})

function copyEmailOnAboutPage() {

	navigator.clipboard.writeText("ljpurcell.dev@gmail.com");

	const path = document.createElementNS("http://www.w3.org/2000/svg", "path");
	path.setAttribute("fill", "#A1A1AA");
	path.setAttribute("d", "M438.6 105.4c12.5 12.5 12.5 32.8 0 45.3l-256 256c-12.5 12.5-32.8 12.5-45.3 0l-128-128c-12.5-12.5-12.5-32.8 0-45.3s32.8-12.5 45.3 0L160 338.7 393.4 105.4c12.5-12.5 32.8-12.5 45.3 0z");

	const tick = document.createElementNS("http://www.w3.org/2000/svg", "svg");
	tick.setAttribute("id", "email-tick");
	tick.setAttribute("class", "ml-2 inline");
	tick.setAttribute("viewBox", "0 0 448 512");
	tick.setAttribute("height", 16);

	tick.appendChild(path);

	const clipboard = document.getElementById("email-clipboard");
	clipboard.replaceWith(tick);
}
