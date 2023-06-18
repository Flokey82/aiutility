# aiutility
Very basic implementation of an utility based AI

The components of the AI are:
- Reasoner
- Action
- Consideration

The reasoner is the main component that decides what to do. It has a list of actions that it can choose from. Each action has a list of considerations that it uses to decide if it should perform that action. Each consideration generates an appraisal (in form of a score) that is used to calculate the utility of the action. The action with the highest utility is then chosen.

Great explanation of utility based AI: https://www.gdcvault.com/play/1012410/Improving-AI-Decision-Modeling-Through