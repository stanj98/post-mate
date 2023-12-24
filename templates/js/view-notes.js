

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

document.addEventListener('DOMContentLoaded', function() {
    function getNotes() {
        return document.getElementsByClassName("note-title");
    }

    document.querySelector(".search-notes-input").addEventListener("change", function(e) {
        var searchText = this.value;
        console.log(getNotes());
        Array.from(getNotes()).forEach(function(note) {
            console.log(note);
            if (searchText.length === 0) {
                note.style.display = "none";
            } else {
                // Do something when there's search text
            }
        });
    });
});

