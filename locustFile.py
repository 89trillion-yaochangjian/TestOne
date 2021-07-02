from locust import HttpUser, task, between

class QuickstartUser(HttpUser):

    @task
    def RarityFindByID(self):
        self.client.get("/RarityFindByID?id=15109")

    @task
    def AtkFindByID(self):
        self.client.get("/AtkFindByID?id=15109")
    @task
    def SoldierFindALLCycUnlock(self):
        self.client.get("/SoldierFindALLCycUnlock?rarity=2&cvc=1000&unlockArena=4")
    @task
    def SoldierFindByCyc(self):
        self.client.get("/SoldierFindByCyc?cvc=1000")
    @task
    def SoldierEachStage(self):
        self.client.get("/SoldierEachStage?unlockArena=4")
