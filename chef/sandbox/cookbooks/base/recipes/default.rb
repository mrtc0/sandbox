#
# Cookbook:: base
# Recipe:: default
#
# Copyright:: 2021, The Authors, All Rights Reserved.

execute 'apt update -qqy'
include_recipe 'base::auditd'
