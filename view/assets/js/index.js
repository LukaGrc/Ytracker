function updateRange(self, id) {
    id.innerHTML = self.value
}
function search(checkbox, creationDate, firstAlbum) {
    let display = true
    Object.entries(checkbox).forEach(v => {
        if (v[1].checked) {
            display = false
        }
    })
    const artists = Array.from(document.getElementsByClassName('artist-card'))
    if (display == true && creationDate.innerHTML == "Sélectionnez" && firstAlbum.innerHTML == "Sélectionnez") {
        artists.forEach(v => {
            v.classList.remove('hidden')
        })
    } else {
        artists.forEach(v => {
            v.classList.add('hidden')
        })
        let ck = false
        Object.entries(checkbox).forEach(v => {
            if (v[1].checked) {
                ck = true
                const checkToDisplay = Array.from(document.getElementsByClassName(`nbMembers-${v[1].value}`))
                checkToDisplay.forEach(val => {
                    val.classList.remove('hidden')
                    if (creationDate.innerHTML != "Sélectionnez") {
                        if (!val.classList.contains(`creationDate-${creationDate.innerHTML}`)) {
                            val.classList.add('hidden')
                        }
                    }
                    if (firstAlbum.innerHTML != "Sélectionnez") {
                        if (!val.classList.contains(`firstAlbumDate-${firstAlbum.innerHTML}`)) {
                            val.classList.add('hidden')
                        }
                    }
                })
            }
        })
        if (!ck) {
            if (creationDate.innerHTML != "Sélectionnez") {
                const creationToDisplay = Array.from(document.getElementsByClassName(`creationDate-${creationDate.innerHTML}`))
                creationToDisplay.forEach(val => {
                    val.classList.remove('hidden')
                    if (firstAlbum.innerHTML != "Sélectionnez") {
                        if (!val.classList.contains(`firstAlbumDate-${firstAlbum.innerHTML}`)) {
                            val.classList.add('hidden')
                        }
                    }
                })
            } else if (firstAlbum.innerHTML != "Sélectionnez" && creationDate.innerHTML == "Sélectionnez") {
                const albumToDisplay = Array.from(document.getElementsByClassName(`firstAlbumDate-${firstAlbum.innerHTML}`))
                albumToDisplay.forEach(val => {
                    val.classList.remove('hidden')
                })
            }
        }
    }
}
function reset(id) {
    id.innerHTML = "Sélectionnez"
    search({ck1, ck2, ck3, ck4, ck5, ck6, ck7}, creationDateValue, firstAlbumDateValue)
}

export { updateRange, search, reset } 