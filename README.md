# Advent of Code

## Get started

### Python

New year:

```bash
mkdir 2025
cp -r 2024/template 2025
uv venv .venv
source .venv/bin/activate
uv pip install advent-of-code-data
uv pip freeze > requirements.txt
```

Paste session cookie into `~/.config/aocd/token`.

Every day:

```bash
cp -r $(date +'%Y')/template $(date +'%Y')/$(date +'%-d')
cd $(date +'%Y')/$(date +'%-d')
source ../.venv/bin/activate
aocd --example
python3 solve.py
```

Paste the example into `example.txt` and uncomment the code block to read from the file.
