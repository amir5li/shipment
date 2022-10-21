package providers

func DetermineSelectedMethods(methods []ValidMethod) (SelectedMethods, error) {
	maxUnit := ^uint(0)
	//maxInt := int(maxUnit >> 1)
	var initialRegularPriority = maxUnit
	var initialPhysicalPriority = maxUnit
	var regularSelectedMethod *ValidMethod
	var physicalSelectedMethod *ValidMethod
	for _, vm := range methods {
		if vm.Physical == true && vm.Priority < initialPhysicalPriority {
			initialPhysicalPriority = vm.Priority
			physicalSelectedMethod = &vm
		}
		if vm.Physical == false && vm.Priority < initialRegularPriority {
			initialRegularPriority = vm.Priority
			regularSelectedMethod = &vm
		}
	}
	var overallSelectedMethod *ValidMethod
	if regularSelectedMethod == nil && physicalSelectedMethod == nil {
		return SelectedMethods{}, NoMethodFound
	}
	if physicalSelectedMethod == nil {
		overallSelectedMethod = regularSelectedMethod
	} else if regularSelectedMethod == nil {
		overallSelectedMethod = physicalSelectedMethod
	} else if regularSelectedMethod.Priority < physicalSelectedMethod.Priority {
		overallSelectedMethod = regularSelectedMethod
	} else {
		overallSelectedMethod = physicalSelectedMethod
	}
	return SelectedMethods{
		Regular:  regularSelectedMethod,
		Physical: physicalSelectedMethod,
		Overall:  overallSelectedMethod,
	}, nil
}
