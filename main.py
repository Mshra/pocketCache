import typer

app = typer.Typer()


@app.command()
def main():
    if not typer.confirm("Are you sure you want to delete it?"):
        print("Not deleting")
        raise typer.Abort()
    print("deleting it!")


if __name__ == "__main__":
    app()
