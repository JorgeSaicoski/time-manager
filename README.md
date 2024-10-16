# Time Manager Application

## Overview

This **Time Manager** application is designed to help employees and freelancers track their working hours, breaks, and projects, providing a detailed breakdown of how much time is invested in each project. Additionally, it offers cost tracking to calculate profits based on a custom hourly rate. The application is built using **Wails** for the backend and **Svelte** for the frontend. The primary goal is to help users understand their work patterns, negotiate salary increases, and calculate project profitability.

## Features

- **Start and Finish Workday**: Log work hours for each day.
- **Break and BRB (Paid Break)**: Track both unpaid and paid break times.
- **Project Management**: Create projects and associate tasks with each work session.
- **Task Tracking**: Log the time spent on each task within a project.
- **Hourly Rate & Profit Calculation**: Calculate how much time has been spent on each project and determine whether a project is profitable based on the set hourly rate.
- **Time Analytics**: Get insights into how much time is spent working each day or on each project.

## Technical Details

- **Backend**: Written in **Golang** using **Wails** framework, with a SQLite database.
- **Frontend**: Built using **Svelte**.
- **Database**: SQLite for local data storage.
- **Deployment**: Planned release on **FlatHub**.

## Installation

### Prerequisites
- **Golang** (Version 1.17+)
- **Wails** (Version 2.0+)
- **Node.js** (For building Svelte frontend)
- **SQLite** (For database management)

### Steps to Run Locally

1. Clone the repository:

   ```bash
   git clone https://github.com/your-repo/time-manager.git
   cd time-manager
   ```

2. Install dependencies:

   ```bash
   npm install
   go mod tidy
   ```

3. Run the application locally:

   ```bash
   wails dev
   ```

4. To build for production:

   ```bash
   wails build
   ```


### Build and Run Locally

To build the project for local use on your computer, follow these steps. This will work on **Linux**, **macOS**, and **Windows**.

1. **Build the frontend**:

   ```bash
   cd frontend
   npm run build
   ```

   This will generate the necessary frontend assets in the `frontend/dist` directory.

2. **Build the Wails project**:

   Navigate back to the root directory of your project and run:

   ```bash
   wails build
   ```

   This command will compile the Go backend and bundle the Svelte frontend into a platform-specific executable.

3. **Run the application**:

   After building the project, you can run the generated executable as follows:

   - **Linux/macOS**: 
     Run the executable from the terminal:
     ```bash
     ./build/bin/time-manager
     ```

   - **Windows**:
     Run the `.exe` file using **PowerShell**, **Command Prompt**, or **Git Bash**:
     ```bash
     .\build\bin\time-manager.exe
     ```


## Usage

### Start a Day

- Upon launching the application, the user can start their workday by clicking the **Start Day** button. The app checks if there’s an unfinished day and prompts the user accordingly.

### Taking Breaks and BRBs

- Users can initiate **Breaks** or **BRBs** (paid breaks) using the dedicated buttons. The app will track the duration of each break and log it accordingly.

### Project Management

- Users can create new projects by entering the project name. Each work session can be associated with a specific project to track time spent on it.

### Work Time Tracking

- As the user works on different tasks, they can log their work time, break times, and check summaries of the time spent on each project.

### Profit Calculation

- Users can enter their hourly rate to calculate how much profit or loss they are making on each project based on the time tracked.

## Roadmap

- **Task Management**: Adding the ability to create and manage specific tasks within each project.
- **Integrations**: Integrating with cloud platforms for backups and cross-device synchronization.
- **Enhanced Analytics**: Offering deeper insights into daily, weekly, or monthly work patterns.

## License

This project is licensed under the MIT License.

## Contributions

Contributions are welcome! Please feel free to submit a pull request or open an issue for any bugs or feature requests.