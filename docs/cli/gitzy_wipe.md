
# gitzy wipe

Irreversibly delete a profile (double-confirm!).

---

## Usage

```bash
gitzy wipe <profile> [flags]
```

This command prompts you twice for confirmation before deleting the profile directory and removing it from the index. This action cannot be undone.

---

## Examples

- Wipe the work profile:
  ```bash
  gitzy wipe work
  ```

---

## Options

- `-h, --help`   Show help for wipe
- `--docs`       Show embedded documentation

---

**Safety Note:** Wiping a profile is permanent. Ensure you have backups if needed. For backup and restore info, see [gitzy switch](gitzy_switch.md) and [gitzy doctor](gitzy_doctor.md).
