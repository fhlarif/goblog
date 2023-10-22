# Project Index: /path/to/project/main.go

## Steps

1. `cd` into the project.

2. Ensure the `env.yaml` is present. Copy it from `app/env/env.yaml.example` and edit the content if necessary.

```bash
cp app/env/env.yaml.example app/env/env.yaml
```

3. To serve the app, run the go script using the following command. This will start the Go development server and make the app available at `http://localhost:8000`. The port is based on the `env.yaml` file.

```bash
go run main.go serve
```

4. To run Tailwind's watch command, make sure you have installed Tailwind and its dependencies.

```bash
pnpm i
```

5. Then, run the following command. This will watch for changes in the source CSS file and automatically rebuild the output file.

```bash
npx tailwindcss -i ./app/resources/assets/css/main.css -o ./app/resources/assets/dist/output.css --build
```
