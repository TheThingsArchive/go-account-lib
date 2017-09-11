# Contributing

There are many ways you can contribute to The Things Network.

- Share your work on the Labs section of our website
- Participate in topics on our Forum
- Talk with others on our Slack (request invite)
- Contribute to our open source projects on github

## Submitting issues

If something is wrong or missing, we want to know about it. Please submit an issue on Github explaining what exactly is the problem. Give as many details as possible. If you can (and want to) fix the issue, please tell us in the issue.

We tried to help you structure your issues by supplying an issue template. Please follow it.

## Contributing pull requests

We warmly welcome your pull requests. Be sure to follow some simple guidelines so that we can quickly accept your contributions.

**Of course follow the golang guidelines**

- Do not use shortened form anywhere.
- Type-scoped functions. Functions do one thing on one type.
- Type functions names and requirement:
  - Types collection functions should have like: `ListTypes`, `StreamTypes`.
  - Type must have the these function: `RegisterType`, `EditType`, `DeleteType`, `GetType`.
  - Type attribute functions should be under this form: `AddTypeAttribute`, `RemoveTypeAttribute`, `ListTypeAttribute`, `GetTypeAttribute`.
- Other functions form are allowed as long as they are justified and start with verb and refer the type, ex: `VerbTypeDoSmth`.
- Documents your codes. You are bringing modification to a library that is used in numerous open project.
- Write tests.
- Run `make test vet lint` (or separately) before pushing.
- Write good commit messages.
- Sign our CLA.
