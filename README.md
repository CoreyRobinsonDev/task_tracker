# task_tracker
[project](https://roadmap.sh/projects/task-tracker)

## Requirements
User should be able to:
[] Add, Update, and Delete tasks
[] Mark a task as in progress or done
[] List all tasks
[] List all tasks that are done
[] List all task that are not done
[] List all tasks that are in progresso

## Example
```bash
# Adding a new task
task_tracker add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
task_tracker update 1 "Buy groceries and cook dinner"
task_tracker delete 1

# Marking a task as in progress or done
task_tracker mark in-progress 1
task_tracker mark done 1

# Listing all tasks
task_tracker list

# Listing tasks by status
task_tracker list done
task_tracker list todo
task_tracker list in-progress
```

## Task Properties
Each task has the following properties:
- **id**: A unique identifier for the task
- **description**: A short description of the task
- **status**: The status of the task (*todo*, *in-progress*, *done*)
- **createdAt**: The date and time when the task was created
- **updatedAt**: The data and time when the task was last updated
