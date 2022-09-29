<template>
  <div class="container">
    <h2 className="title_airdrop">Airdrop</h2>
    <h1 className="sub_title_airdrop">Your progress</h1>
    <div class="flex flex-wrap p-4">
      <SpProgressbar size="large" val="60" text="60%" bar-border-radius="20" text-align="right">
     </SpProgressbar>
    </div>
    <h2 className="title_airdrop">My Missions</h2>

    <div className="mission-group  d-flex align-items-right justify-content-between"> 
      <div class="col-md-6"> 
        <div>
          <p className="mission-title"> Mission #1 </p>
        </div>
        <div>
          <p className="mission-content">initual claim (25%)</p>
        </div>
      </div>
      <div>
         <Button class="Button">
           Claim
         </Button>
       </div>
    </div>
    <div className="mission-group  d-flex align-items-right justify-content-between"> 
      <div class="col-md-6"> 
        <div>
          <p className="mission-title"> Mission #2 </p>
        </div>
        <div>
          <p className="mission-content">stake the claimed in 1) to unlock 25% more (25%)</p>
        </div>
      </div>
      <div>
         <Button class="Button">
           Stake
         </Button>
       </div>
    </div>
    <div className="mission-group  d-flex align-items-right justify-content-between"> 
      <div class="col-md-6"> 
        <div>
          <p className="mission-title"> Mission #3 </p>
        </div>
        <div>
          <p className="mission-content">vote with staked balance for a proposal (25%)</p>
        </div>
      </div>
      <div>
         <Button class="Button">
           Vote
         </Button>
       </div>
    </div>
    <div className="mission-group  d-flex align-items-right justify-content-between"> 
      <div class="col-md-6"> 
        <div>
          <p className="mission-title"> Mission #4 </p>
        </div>
        <div>
          <p className="mission-content">perform a swap (25%)</p>
        </div>
      </div>
      <div>
         <Button class="Button">
           Swap
         </Button>
       </div>
    </div>
  </div>
</template>

<!--***************type script section start *************************** -->
<script lang="ts">
import { defineComponent } from 'vue'
import { useStore } from 'vuex'

import  SpProgressbar from '../SpProgressbar'

export default defineComponent({
  name: 'SpAirmission',

  components: {
    SpProgressbar
  },
 
})
</script>

<!--*************** CSS section start *************************** -->
<style lang="scss" scoped>
$base-color: rgba(0, 0, 0, 0.03);
$animation-duration: 1.6s;
$shine-color: rgba(0, 0, 0, 0.06);
$avatar-offset: 32 + 16;
.assets-header {
    display: flex;
    flex-wrap: wrap;
    width: 100%;
    
}
.title_airdrop {
  font-family: Inter, serif;
  font-style: normal;
  font-weight: 700;
  font-size: 28px;
  line-height: 127%;
  /* identical to box height, or 36px */

  letter-spacing: -0.02em;
  font-feature-settings: 'zero';

  color: #ffffff;
  margin-top: 2%;
}
.sub_title_airdrop {
  font-family: Inter, serif;
  font-style: normal;
  font-size: 18px;
  margin-top: 2%;
  line-height: 127%;
  /* identical to box height, or 36px */
  padding: 10px;
  font-feature-settings: 'zero';
  color: #ffffff;

}
.mission-title {
  font-family: Inter, serif;
  font-style: normal;
  font-weight: 700;
  font-size: 14px;
  /* identical to box height, or 36px */
  margin-top:15px;
  font-feature-settings: 'zero';
  color: #ffffff;
  margin-top: 0;
}
.mission-content {
  font-family: Inter, serif;
  font-style: normal;
  font-weight: 700;
  font-size: 18px;
  line-height: 127%;
  /* identical to box height, or 36px */
  
  font-feature-settings: 'zero';
  color: #ffffff;
  
}


.mission-group {
  display: flex;
  justify-content: space-between;
  margin-top: 2%;
  border: rgba(151, 130, 130, 0.1);
  border-radius: 10px;
  border: 1px solid grey;
  padding-top: 25px;
}
.claim-button-style {
  display: flex;
  justify-content: space-between;
  height: auto;
  width: auto;
  font-size:14px;
  background: #BE185D;
  border-radius: 5px;
  padding: 5px;
}
.Button {
  background:#BE185D;
  width: 80px;
  border: none;
  color: #fff;
  cursor: pointer;
  font-size: 14px;
  margin-top: 5px;
  margin-right:20px;
  padding: 12px 20px;
  border-radius:15px;
  border-color:#BE185D;
  transition: background .2s ease-in-out;
  &:hover {
    background: darken(#d11e69, 20);
  }
  &:disabled {
    opacity: 0.5;
  }
}
</style>
