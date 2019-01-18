const { context } = require("./tools");

module.exports = () => {
  context.slackMessage({
    text:
      "[inextensodigital/invoicing.app:0.1.1] Your release are ready to deploy !",
    attachments: [
      {
        actions: [
          {
            type: "button",
            text: "Deploy 🚀",
            url:
              "https://auto-deploy.inextenso.io/deploy?owner=inextensodigital&repo=actions-testing&deploy=124063846&sign=62fc1e52b6308af303800f94d8d9c00479cf5c08e9f1b69a41fc98c1707086128c76e75b0acb1fd1c80b7d0b7681216dc7acd112782ac17ac37170c1d969a567c3b33493746cbbd3d3ba88088860135b52a77ffdee773d5fb89b3cb0cf22b3061dc859cdb8879a7882fe09326c6dd119399aa171191587dce1bfdd1ee0ea3883d251f18b2813dcfbee80db1130ae6856b716544f96f3c6c8d315a98e5f320c90ace2a76ca4069417af46eae7bd78026bc39fa53e107fb2dd41278f0bbdb54c780bb868f4e7b66c6b3d6da80073f334ffd68ec96482c4c8938a04fc59ca91c20e5910af533731a4698b99547962de2ac161b56a55693da6d5c8392a12f0588306",
            style: "danger",
            confirm: {
              title: "Are you sure?",
              text: "ptit mami?",
              ok_text: "Yes",
              dismiss_text: "No"
            }
          }
        ]
      }
    ]
  });
};
