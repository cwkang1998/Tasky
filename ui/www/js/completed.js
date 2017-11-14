function initCompleted(tasks) {
    //Set placeholder invisible and content visible
    let ph = document.getElementById("placeholder-completed");
    ph.style.display = "none";

    let content = document.getElementsByClassName("completed");
    for (let i = 0; i < content.length; i += 1) {
        content[i].style.display = "block";
    }

    //Population of the entries
    let fragment = document.createDocumentFragment();
    for (let i = 0; i < tasks.length; i++) {
        console.log(tasks[i]);
        let taskElements = createTask(tasks[i], "completed");
        fragment.appendChild(taskElements);
    }
    document.getElementById("completedTasks").appendChild(fragment);
}

function showEmptyCompleted() {
    let ph = document.getElementById("placeholder-completed");
    ph.style.display = "inline";

    let content = document.getElementsByClassName("completed");
    for (let i = 0; i < content.length; i += 1) {
        content[i].style.display = "none";
    }
}

function attachShowOptions(id) {
    return function () {
        $('#alert-pop-up')
            .append("<div id='alert-div' class='fade in alert alert-dismissible alert-warning'>" +
            "<span class='close' data-dismiss='alert' aria-label='close'>&times;</span>" +
            "<h4> What do you want to do with this task entry?</h4>" +
            "<div><button id='resetTaskBtn' class='btn btn-default' style='margin-right:14px;'>Reset as active task</button>" +
            "<button id='deleteTaskBtn' class='btn btn-danger btn-margin' style='margin-right:14px;'>Delete Task</button></div>" +
            "</div>");
        document.getElementById("resetTaskBtn").onclick = function () {
            setPending(id);
            $("#alert-div").alert("close");
        };
        document.getElementById("deleteTaskBtn").onclick = function () {
            deleteTask(id);
            $("#alert-div").alert("close");
        };
    };
}


function setPending(id) {
    makeGetRequest("GET", "/setTskStatus",
        function () {
            reloadAll();
            showAlert("Task Reset to Pending.", "alert-success");
        }
        , function () {
            showAlert("Action Failed. Please try again later", "alert-danger");
        }
        , "?task_id=" + encodeURIComponent(id) + "&status=0");

}

function deleteTask(id) {
    makeGetRequest("GET", "/delTsk",
        function () {
            reloadCompleted();
            showAlert("Task Deleted.", "alert-success");
        }
        , function () {
            showAlert("Action Failed. Please try again later", "alert-danger");
        }
        , "?task_id=" + encodeURIComponent(id));
}