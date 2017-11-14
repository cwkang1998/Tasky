
// Toggle button setup for add new task button
document.getElementById("addNewTasks").onclick = function () {
    let text = document.getElementById("newTasks");
    let input = document.getElementById("taskInput");

    // Receive enter key, if enter key pressed then save to db
    input.onkeypress = function (e) {
        if (e.keyCode == 13) {

            let taskEntry = input.value;
            if (taskEntry.length > 0) {
                text.style.display = "none";
                input.value = "";
                input.placeholder = "";
                saveTasks(taskEntry);
            } else {
                input.value = "";
                input.placeholder = "Please Add A Task with more than 1 character :)";
            }

        }
    };

    if (text.style.display == "none") {
        text.style.display = "block";
        input.focus();

    } else {
        text.style.display = "none";
        input.placeholder = "";
        input.value = "";
    }
};

reloadAll();