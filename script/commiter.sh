#!/bin/bash

# Get the current date in the format YYYY-MM-DD
current_date=$(date +"%Y-%m-%d")

# Get the current date and time 
current_date_time=$(date +"%Y-%m-%d %H:%M:%S")

# Define the branch name based on the current date
branch_name="daily-branch-$current_date"

# Check if there are changes in the working directory
if [ -z "$(git status --porcelain)" ]; then
    echo "No changes to commit."
else
    # Check if the branch already exists
    if git show-ref --quiet "refs/heads/$branch_name"; then
        echo "Branch '$branch_name' already exists. Exiting."
        exit 0
    fi

    # Create a new branch for the current day
    git checkout -b "$branch_name"

    # Commit changes with a message indicating the date
    git add .
    git commit -m "Changes for $current_date_time"
    
    # Push the changes to the remote repository
    git push origin "$branch_name"
    
    echo "Changes committed and pushed to $branch_name"

    # Checkout back to the original branch
    git checkout -
    
    # Reset the last commit (if one was made)
    if [ -n "$(git rev-parse --verify HEAD)" ]; then
        git merge "$branch_name"
        git reset HEAD~1
        echo "Last commit reset."
    fi
fi

