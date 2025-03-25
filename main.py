import click


@click.group()
@click.option(
    "--config",
    "-c",
    default="config.yaml",
    type=click.Path(exists=True, readable=True),
    help="Path to the configuration file (YAML format).",
)
@click.pass_context
def cli(ctx: click.Context, config: str) -> None:
    """
    Infrastructure Automation CLI Tool

    A command-line interface for managing infrastructure provisioning
    and teardown operations.
    """
    ctx.ensure_object(dict)


@cli.command()
@click.pass_context
def create(ctx: click.Context) -> None:
    """
    Provision infrastructure based on configuration

    Creates infrastructure resources as specified in the config file.
    """
    raise NotImplementedError


@cli.command()
@click.pass_context
def delete(ctx: click.Context) -> None:
    """
    Destroy provisioned infrastructure

    Removes all infrastructure resources created by the create command.
    """
    raise NotImplementedError


if __name__ == "__main__":
    cli(obj={})
