

//Delete note functionality

async function deleteNote() {
    try {
        const deleteAttr = document.querySelector(".icon");
        const noteID = deleteAttr.dataset.imgid;
        const noteContainer = document.getElementById("note-"+ noteID);
        const inputContainer = document.querySelector(".search-notes-input");
        var arrayLength = document.getElementById("arrLength");
        await axios.delete(ROOT_API + "api/notes/" + noteID)
        .then(response => {
            if (response.status == 200) {
                noteContainer.remove();
                arrayLength.value -= 1;
                if (arrayLength.value == 0) {
                    inputContainer.remove();
                    const emptyMsgContainer = document.createElement('div');
                    emptyMsgContainer.innerHTML = `
                        <p> You haven't created any note yet. <a href="/create-note">Create a note now</a></p>
                    `;
                    document.querySelector('.main-section').appendChild(emptyMsgContainer);
                }
            }
        });
    }
    catch(errors) {
        console.log(errors);
    }
}

//Search note functionality
function getNotes() {
    return document.getElementsByClassName("note-title-link");
}

function searchNote(e) {
    var searchText = e.value;
    Array.from(getNotes()).forEach(function(note) {
        var noteText = note.text;
        if (noteText.search(searchText) != 0) {
            note.parentElement.parentElement.style.display = "none";
        } else {
            note.parentElement.parentElement.style.display = "";
        }
    });
}

