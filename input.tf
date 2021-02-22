resource "aws_sqs_queue" "sqs_2" {
  name                              = "terraform-example-queue"
  kms_master_key_id                 = "alias/aws/sqs"
  kms_data_key_reuse_period_seconds = 300
}
