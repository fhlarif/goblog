# FILEPATH: /path/to/project/main.go

"""
This script serves the app and runs Tailwind's watch command.

To serve the app, run the script using the command:

```bash
go run main.go serve
```

This will start the Go development server and make the app available at http://localhost:8000.

To run Tailwind's watch command, make sure you have installed Tailwind and its dependencies.

```bash
pnpm i
```

Then, run the command:

```bash
npx tailwindcss -i ./app/resources/assets/css/main.css -o ./app/resources/assets/dist/output.css --build
```

This will watch for changes in the source CSS file and automatically rebuild the output file.
"""
