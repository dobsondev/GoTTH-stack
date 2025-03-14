# GoTTH Stack

- Go
- Templ
- Tailwindcss
- HTMX

## Running Locally

Everything can be run using `docker compose`:

```bash
docker compose up -d
```

This will automatically hot reload with (air)[https://github.com/air-verse/air], utilize a proxy for automatic browser refreshing and watch and recompile tailwind styles as needed. The application can be accessed at:

- http://localhost:8080/

## Manual Running

To install dependencies:

```bash
bun install
```

To compile our Tailwind styles:

```bash
bunx @tailwindcss/cli -i ./styles/input.css -o ./styles/style.css --minify --watch=always
```

Using (air)[https://github.com/air-verse/air] to hot reload and live reload our application:

```bash
air -c .air.toml
```