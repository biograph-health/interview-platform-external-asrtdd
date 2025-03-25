# Biograph Platform Technical Interview

From the Biograph engineering team, thank you for your interest in a platform engineering role. We are excited to meet you!

During one of your interviews, you will pair with an engineer on the team to add some functionality to this pretend codebase. Our goal is to learn how you approach work that is similar in nature to what you may do at Biograph on a day-to-day basis.

At Biograph, we are primarily interested in building useful and delightful products. That's what we are striving to assess here. We are not interested in quizzing you on algorithms, and there are no trick questions. You'll be asked to make a series of decisions without clear "right" or "wrong" answers. We hope this allows you to most effectively demonstrate how you collaborate and think through technical problems.

We want you to use whatever coding environment you are most comfortable with, so we invite you to bring your own laptop to the interview. We'll ask you to share your screen while you are writing code. Read through the rest of this document for a brief introduction to the codebase and instructions on running and testing the application.

## Scenario

Your company aims to simplify cloud infrastructure provisioning by developing a CLI tool. The tool should automate the
creation and deletion S3 resources using Terraform.

## Prerequisites

- Docker Compose or equivalent

## Requirements

The CLI should support the following commands:

- create - Provisions a S3 bucket and object.
- delete - Destroys the infrastructure.

## Implementation Details

### Language

- Preferred languages include Python (with click or argparse) and Go (with cobra), but candidates are free to use other
  languages if they can demonstrate effective Terraform integration and usability.

### Terraform Integration

- The CLI should invoke Terraform commands (init, apply, destroy) using a system process (e.g., Python’s subprocess,
  Go’s os/exec, or equivalent in other languages).
- Terraform wrappers (e.g., terraform-exec) may be used but are not required.

### Configuration Management

- Use a config.yaml file for user-defined settings and resource specifications.
- Example configuration file:

```yaml
minio:
  server: 127.0.0.1:9000
  user: developer
  password: strong_password

bucket_name: platform-interview
file_content: Lorem ipsum odor amet, consectetuer adipiscing elit.
```

## State Management

- Store Terraform state locally by default.

## Logging & Error Handling

- Implement structured logging to provide meaningful output to users.
- Ensure Terraform failures are handled gracefully with proper error messages.
- Validate configuration inputs before executing commands.

## Example CLI Usage

```bash
# Provision infrastructure
$ infra-cli create --config config.yaml

# Destroy infrastructure
$ infra-cli delete --config config.yaml
```

## Evaluation Criteria

Candidates will be assessed based on the following:

- Code Quality - Clean, modular, and well-structured implementation.
- Terraform Integration - Effective interaction with Terraform.
- Error Handling & Logging - Robust handling of failures and informative logs.
- Automation & Usability - Ease of use, flexibility, and maintainability.
- Security Best Practices - Proper management of credentials, permissions, and Terraform state.
