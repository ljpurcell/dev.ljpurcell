dev.ljpurcell — Site Reference
==============================

Stack & structure
-----------------

- Plain HTML/CSS — no framework, no build step, no package manager, no CI
- Three files: `index.html`, `style.css`, `netlify.toml`
- Deployed on Netlify at **ljpurcell.com**; changes deploy directly on push
- Test in-browser (open `index.html`) — no dev server needed

Background
----------

- Redesigned from a Go/Tailwind server (DigitalOcean droplet) to static HTML/CSS on Netlify — a single static page never justified a server, and Netlify gives HTTPS, CDN, and compression for free
- The senior-portfolio redesign (content: work history, upstream/leverage tagline; theme matching the Hot & Dongerous palette) was on branch `redesign/senior-portfolio` as of May 2026
- At that point the old Go server was still live on the droplet — verify the current deploy state rather than assuming
