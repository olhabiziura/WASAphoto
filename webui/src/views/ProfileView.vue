<script>
// get user profile
export default {
    data: function() {
        return {
            errormsg: null,

            // getUserProfile
            username: "",
            photosCount: 0,
            followersCount: 0,
            followingCount: 0,
            isOwner: false,
            doIFollowUser: false,

            isInMyBannedList: false,
            amIBanned: false,

            // getPhotosList
            photosList: [],

            // getFollowersList
            followerList: [],

            // getFollowingsList
            followingList: [],
            
            banList: [],

            userExists: false,
            user_id: 0,
        }
    },
    watch: {
        // property to watch
        pathUsername(newUName, oldUName) {
            if (newUName !== oldUName) {
                this.getUserProfile()
            }
        }
    },
    computed: {
        pathUsername() {
            return this.$route.params.username
        },
    },
    methods: {
        async getUserProfile() {
            if (this.$route.params.username === undefined) {
                return
            }
            try {
                let username = this.$route.params.username;
                let response = await this.$axios({
                    method: 'post',
                    url: `/session`,
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    data: {
                        username: username
                    }
                });
                this.user_id = response.data.user_id;
                console.log(sessionStorage.getItem('token'))
                response = await this.$axios.get(`/profile/${this.user_id}`, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                let profile = response.data;
                this.username = profile.Username;
                this.banList = profile.banList || []
                this.amIBanned = this.banList.some(user =>user === sessionStorage.getItem('token'));
                if (this.amIBanned){
                    await this.$axios.delete(`/follow/${this.user_id}`, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                    this.getUserProfile();
                }

                
                console.log("before")
                response = await this.$axios.get(`/ban/${this.user_id}`, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                this.isInMyBannedList = response.data.isBanned;
                console.log("after")
                console.log(response)
                


                
            
        

                // Sort photos by date in descending order
                this.photosList = profile.photoList.sort((a, b) => new Date(b.Date) - new Date(a.Date));
               
                this.followerList = profile.followerList || []
                this.doIFollowUser = this.followerList.some(user => user.UserID === sessionStorage.getItem('token'));

                this.followingList = profile.followingList || []
                if (sessionStorage.getItem('username') === username.toLowerCase()) {
                    this.isOwner = true;
                }
                
                if (profile.photoList != null) {
                    this.photosCount = profile.photoList.length;
                } else {
                    this.photosCount = 0;
                }

                if (profile.followerList != null) {
                    this.followersCount = profile.followerList.length;
                    for (let i = 0; i < profile.followerList.length; i++) {
                        if (profile.followerList[i].user_id === sessionStorage.getItem('token')) {
                            this.doIFollowUser = true;
                            break;
                        }
                    }
                } else {
                    this.followersCount = 0;
                }
                if (profile.followingList != null) {
                    this.followingCount = profile.followingList.length;
                } else {
                    this.followingCount = 0;
                }
                if (profile.user_id === sessionStorage.getItem('token')) {
                    this.isOwner = true;
                }

                this.userExists = true;
                if (!this.isInMyBannedList && !this.amIBanned) {
                    await this.getPhotosList();
                    this.getFollowersList();
                    this.getFollowingsList();
                }
            } catch (error) {
                const status = error.response.status;
                const reason = error.response.data;
                this.errormsg = 'Status' + status + ': ' + reason;
            }
        },
        async followBtn() {
            try {
                if (this.doIFollowUser) { 
                     // DELETE /unfollow/{uid}
                    await this.$axios.delete(`/follow/${this.user_id}`, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                    this.getUserProfile();
                } else {
                    // POST /follow/{uid}
                    await this.$axios.post(`/follow/${this.user_id}`, null, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                    this.getUserProfile();
                }
                this.doIFollowUser = !this.doIFollowUser
            } catch (error) {
                const status = error.response.status;
                const reason = error.response.data;
                this.errormsg = 'Status' + status + ': ' + reason;
            }
        },
        async banBtn() {
            try {
                if (this.isInMyBannedList) {
                    // DELETE /ban/{uid}
                    await this.$axios.delete(`/ban/${this.user_id}`, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                    this.isInMyBannedList = false;
                    this.getUserProfile();
                } else {
                    // POST /ban/{uid}
                    await this.$axios.post(`/ban/${this.user_id}`, null, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                    this.isInMyBannedList = true;
                    this.getUserProfile();
                }
            } catch (error) {
                const status = error.response.status;
                const reason = error.response.data;
                this.errormsg = 'Status' + status + ': ' + reason;
            }
        },
        async uploadPhoto() {
            try {
                let file = document.getElementById('fileUploader').files[0];
                
                const reader = new FileReader();
                reader.readAsArrayBuffer(file); // stored in result attribute
                let formData = new FormData();
                formData.append('picture', file);
                formData.append('description', 'Your photo description here')
                reader.onload = async () => {
                    // POST /photo
                    let response = await this.$axios.post('/photo', formData, {
                        headers: {
                            'Authorization': `${sessionStorage.getItem('token')}`,
                            'Content-Type': 'multipart/form-data'
                        }
                    });                    
                    this.photosList.unshift(response.data); // at the beginning of the list
                    this.photosCount += 1;
                }
            } catch (error) {
                const status = error.response.status;
                const reason = error.response.data;
                this.errormsg = 'Status' + status + ': ' + reason;
            }
        },

        // on child event
        removePhotoFromList(pid) {
            this.photosList = this.photosList.filter(photo => photo.PictureID != pid);
            this.photosCount -= 1;
        },
        visitUser(username) {
            if (username != this.$route.params.username) {
                this.$router.push(`/profile/${username}`);
            }
        }
    },
    mounted() {
        this.getUserProfile();
      

    }
}
</script>

<template>

    <UserModal
    :modalID="'usersModalFollowers'" 
    :usersList="followerList"
    @visitUser="visitUser"
    />

    <UserModal
    :modalID="'usersModalFollowing'" 
    :usersList="followingList"
    @visitUser="visitUser"
    />


    <div class="container-fluid">
        <div v-if="amIBanned" class="row">
            <div class="col-12 d-flex justify-content-center">
                <h2>Sorry! You have been banned by this user.</h2>
                <button v-if="!isOwner" @click="banBtn" class="btn btn-danger ms-2">
                                        {{isInMyBannedList ? "Unban" : "Ban"}}
                </button>
            </div>
        </div>
        
        <div v-else-if="userExists" class="container-fluid">
            <div class="row">
                <div class="col-12 d-flex justify-content-center">
                    <div class="card w-50 container-fluid">
                        <div class="row">
                            <div class="col">
                                <div class="card-body d-flex justify-content-between align-items-center">
                                    <h5 class="card-title p-0 me-auto mt-auto">@{{username}}</h5>

                                    <button v-if="!isOwner && !isInMyBannedList" @click="followBtn" class="btn btn-success ms-2">
                                        {{doIFollowUser ? "Unfollow" : "Follow"}}
                                    </button>

                                    <button v-if="!isOwner" @click="banBtn" class="btn btn-danger ms-2">
                                        {{isInMyBannedList ? "Unban" : "Ban"}}
                                    </button>
                                </div>
                            </div>
                        </div>

                        <div v-if="!isInMyBannedList" class="row mt-1 mb-1">
                            <button class="col-4 d-flex justify-content-center btn-foll">
                                <h6 class="ms-3 p-0 ">Posts: {{photosCount}}</h6>
                            </button>
                        
                            <button class="col-4 d-flex justify-content-center btn-foll">
                                <h6 data-bs-toggle="modal" :data-bs-target="'#usersModalFollowers'">
                                    Followers: {{followersCount}}
                                </h6>
                            </button>
                        
                            <button class="col-4 d-flex justify-content-center btn-foll">
                                <h6 data-bs-toggle="modal" :data-bs-target="'#usersModalFollowing'">
                                    Following: {{followingCount}}
                                </h6>
                            </button>
                        </div>
                    </div>
                </div>
            </div>


            <div class="row">
                <div class="container-fluid mt-3">
                    <div class="row ">
                        <div class="col-12 d-flex justify-content-center">
                            <h2>Posts</h2>
                            <input id="fileUploader" type="file" class="profile-file-upload" @change="uploadPhoto" accept=".jpg, .png">
                            <label v-if="isOwner" class="btn my-btn-add-photo ms-2 d-flex align-items-center" for="fileUploader"> Add </label>
                        </div>
                    </div>
                    <div class="row ">
                        <div class="col-3"></div>
                        <div class="col-6">
                            <hr class="border border-dark">
                        </div>
                        <div class="col-3"></div>
                    </div>
                </div>
            </div>

            <div class="row">
                <div class="col">
                    <div v-if="!isInMyBannedList && photosCount > 0">
                        <Photo v-for="photo in photosList"
                            :key="photo.PictureID"
                            :pid="photo.PictureID"
                            :ownerID="photo.OwnerID"
                            :username="photo.Username"
                            :date="photo.Date"
                            :likesListParent="photo.likes"
                            :commentsListParent="photo.Comments"
                            :isOwner="isOwner"
                            :image="photo.image"
                            @removePhoto="removePhotoFromList"
                        />
                    </div>
                    
                    <div v-if="!isInMyBannedList && photosCount == 0" class="mt-5">
                        <h2 class="d-flex justify-content-center" style="color: white;">No posts yet</h2>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style>
.profile-file-upload {
    display: none;
}
.my-btn-add-photo {
    background-color: green;
    border-color: grey;
}
.my-btn-add-photo:hover {
    color: white;
    background-color: green;
    border-color: grey;
}
.btn-foll {
    background-color: transparent;
    border: none;
    padding: 5px;
}
</style>
